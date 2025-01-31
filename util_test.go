// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"bytes"
	"errors"
	"testing"
)

// mockReader is a mock implementation of io.Reader for testing
type mockReader struct {
	data []byte // Data to be read
	err  error  // Error to return during read
}

func (m *mockReader) Read(p []byte) (n int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	n = copy(p, m.data)
	return n, nil
}

// TestRandomBits tests the randomBits function
func TestRandomBits(t *testing.T) {
	originalRander := rander
	defer func() { rander = originalRander }()

	t.Run("fills slice with random data", func(t *testing.T) {
		mockData := []byte{0x01, 0x02, 0x03, 0x04}
		rander = &mockReader{data: mockData}

		b := make([]byte, len(mockData))
		randomBits(b)

		if !bytes.Equal(b, mockData) {
			t.Errorf("expected %v, got %v", mockData, b)
		}
	})

	t.Run("panics on read error", func(t *testing.T) {
		rander = &mockReader{err: errors.New("read error")}

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic, but did not occur")
			} else if errMsg, ok := r.(string); !ok || errMsg != "read error" {
				t.Errorf("expected panic with message 'read error', got %v", r)
			}
		}()

		randomBits(make([]byte, 4))
	})
}
