package uuid

import (
	"sync"
	"testing"
	"time"
)

func TestNewV6WithTime(t *testing.T) {
	testCases := map[string]string{
		"test with current date":                      time.Now().Format(time.RFC3339),                                // now
		"test with past date":                         time.Now().Add(-1 * time.Hour * 24 * 365).Format(time.RFC3339), // 1 year ago
		"test with future date":                       time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339),      // 1 year from now
		"test with different timezone":                "2021-09-01T12:00:00+04:00",
		"test with negative timezone":                 "2021-09-01T12:00:00-12:00",
		"test with future date in different timezone": "2124-09-23T12:43:30+09:00",
	}

	for testName, inputTime := range testCases {
		t.Run(testName, func(t *testing.T) {
			customTime, err := time.Parse(time.RFC3339, inputTime)
			if err != nil {
				t.Errorf("time.Parse returned unexpected error %v", err)
			}
			id, err := NewV6WithTime(&customTime)
			if err != nil {
				t.Errorf("NewV6WithTime returned unexpected error %v", err)
			}

			if id.Version() != 6 {
				t.Errorf("got %d, want version 6", id.Version())
			}
			unixTime := time.Unix(id.Time().UnixTime())
			// Compare the times in UTC format, since the input time might have different timezone,
			// and the result is always in system timezone
			if customTime.UTC().Format(time.RFC3339) != unixTime.UTC().Format(time.RFC3339) {
				t.Errorf("got %s, want %s", unixTime.Format(time.RFC3339), customTime.Format(time.RFC3339))
			}
		})
	}
}

func TestNewV6FromTimeGeneratesUniqueUUIDs(t *testing.T) {
	now := time.Now()
	ids := make([]string, 0)
	runs := 26000

	for i := 0; i < runs; i++ {
		now = now.Add(time.Nanosecond) // Without this line, we can generate only 16384 UUIDs for the same timestamp
		id, err := NewV6WithTime(&now)
		if err != nil {
			t.Errorf("NewV6WithTime returned unexpected error %v", err)
		}
		if id.Version() != 6 {
			t.Errorf("got %d, want version 6", id.Version())
		}

		// Make sure we add only unique values
		if !contains(t, ids, id.String()) {
			ids = append(ids, id.String())
		}
	}

	// Check we added all the UIDs
	if len(ids) != runs {
		t.Errorf("got %d UUIDs, want %d", len(ids), runs)
	}
}

func TestNewV6WithTimeConcurrentUnique(t *testing.T) {
	// NewV6WithTime must take the clock-sequence lock so concurrent generation
	// for the same timestamp stays race-free and keeps producing unique values.
	// For a fixed timestamp the UUID varies only by the clock sequence, so up
	// to 16384 calls must all be unique. Run with -race to also surface the
	// underlying data race directly.
	fixed := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	const goroutines = 8
	const perGoroutine = 2000 // 16000 total < 16384 clock-sequence values

	var wg sync.WaitGroup
	var mu sync.Mutex
	seen := make(map[UUID]struct{}, goroutines*perGoroutine)
	dups := 0

	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perGoroutine; i++ {
				id, err := NewV6WithTime(&fixed)
				if err != nil {
					t.Errorf("NewV6WithTime returned unexpected error %v", err)
					return
				}
				mu.Lock()
				if _, ok := seen[id]; ok {
					dups++
				} else {
					seen[id] = struct{}{}
				}
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	if dups != 0 {
		t.Errorf("got %d duplicate V6 UUIDs from concurrent NewV6WithTime calls", dups)
	}
}

func BenchmarkNewV6WithTime(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			now := time.Now()
			_, err := NewV6WithTime(&now)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func contains(t *testing.T, arr []string, str string) bool {
	t.Helper()

	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
