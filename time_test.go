package uuid

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	now := time.Now()
	tt := map[string]struct {
		input        func() *time.Time
		expectedTime int64
	}{
		"it should return the current time": {
			input: func() *time.Time {
				return nil
			},
			expectedTime: now.Unix(),
		},
		"it should return the provided time": {
			input: func() *time.Time {
				parsed, err := time.Parse(time.RFC3339, "2024-10-15T09:32:23Z")
				if err != nil {
					t.Errorf("timeParse unexpected error: %v", err)
				}
				return &parsed
			},
			expectedTime: 1728984743,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			result, _, err := getTime(tc.input())
			if err != nil {
				t.Errorf("getTime unexpected error: %v", err)
			}
			sec, _ := result.UnixTime()
			if sec != tc.expectedTime {
				t.Errorf("expected %v, got %v", tc.expectedTime, result)
			}
		})
	}
}

func TestClockSequenceInitialization(t *testing.T) {
	// Backup and reset global state
	timeMu.Lock()
	origSeq := clockSeq
	clockSeq = 0 // Force initialization
	timeMu.Unlock()
	defer func() {
		timeMu.Lock()
		clockSeq = origSeq
		timeMu.Unlock()
	}()

	// First call initializes the sequence
	seq := ClockSequence()
	if seq == 0 {
		t.Error("ClockSequence() should not return 0 after initialization")
	}

	// Ensure the sequence is within 14-bit mask
	if seq&0xc000 != 0 {
		t.Errorf("ClockSequence() = %x, expected 14-bit value (mask 0x3fff)", seq)
	}

	// Subsequent call should return the same sequence
	seq2 := ClockSequence()
	if seq2 != seq {
		t.Errorf("Expected sequence %x, got %x", seq, seq2)
	}
}

func TestTime(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		version int
		want    Time
	}{
		{
			name:    "Version 6",
			uuid:    UUID{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0x6D, 0xEF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			version: 6,
			want:    Time((0x12345678 << 28) | (0x9ABC << 12) | 0xDEF),
		},
		{
			name:    "Version 7",
			uuid:    UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			version: 7,
			want:    Time(g1582ns100),
		},
		{
			name:    "Default Version",
			uuid:    UUID{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0x1D, 0xEF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			version: 1,
			want:    Time(int64(0x12345678) | (int64(0x9ABC) << 32) | (int64(0xDEF) << 48)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.Time(); got != tt.want {
				t.Errorf("UUID.Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClockSequence(t *testing.T) {
	uuid := UUID{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	expected := 0x1234 & 0x3fff
	if got := uuid.ClockSequence(); got != expected {
		t.Errorf("ClockSequence() = %x, want %x", got, expected)
	}
}
