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
