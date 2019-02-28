package uuid

import (
	"math/rand"
	"testing"
	"time"
)

func TestUuidSources(t *testing.T) {
	
	// Two identical sources, should give same sequence
	currentTime := time.Now().UnixNano()
	uuidSourceA := NewSource(rand.New(rand.NewSource(currentTime)))
	uuidSourceB := NewSource(rand.New(rand.NewSource(currentTime)))
	
	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 != uuid2 {
			t.Errorf("expected duplicates, got %q and %q", uuid1, uuid2)
		}
	}
	
	// Set rander with nil, each source will be random
	uuidSourceA.SetRand(nil)
	uuidSourceB.SetRand(nil)

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q", uuid1)
		}
	}
	
	// Set rander to rand source with same seed, should give same sequence
	uuidSourceA.SetRand(rand.New(rand.NewSource(123)))
	uuidSourceB.SetRand(rand.New(rand.NewSource(123)))

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 != uuid2 {
			t.Errorf("expected duplicates, got %q and %q", uuid1, uuid2)
		}
	}

	// Set rander to rand source with different seeds, should not give same sequence
	uuidSourceA.SetRand(rand.New(rand.NewSource(456)))
	uuidSourceB.SetRand(rand.New(rand.NewSource(789)))

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q", uuid1)
		}
	}

}