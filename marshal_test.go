// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"encoding/json"
	"testing"
)

func TestMarshalText(t *testing.T) {
	u := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	expected := "01234567-89ab-cdef-fedc-ba9876543210"

	data, err := u.MarshalText()
	if err != nil {
		t.Fatalf("MarshalText failed: %v", err)
	}

	if string(data) != expected {
		t.Errorf("MarshalText() = %s; want %s", data, expected)
	}
}

func TestUnmarshalText(t *testing.T) {
	var u UUID
	input := []byte("01234567-89ab-cdef-fedc-ba9876543210")
	expected := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}

	err := u.UnmarshalText(input)
	if err != nil {
		t.Fatalf("UnmarshalText failed: %v", err)
	}

	if u != expected {
		t.Errorf("UnmarshalText() = %v; want %v", u, expected)
	}
}

func TestUnmarshalTextInvalidFormat(t *testing.T) {
	var u UUID
	input := []byte("invalid-uuid-format")

	err := u.UnmarshalText(input)
	if err == nil {
		t.Fatalf("UnmarshalText should have failed for invalid format")
	}
}

func TestMarshalBinary(t *testing.T) {
	u := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	expected := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}

	data, err := u.MarshalBinary()
	if err != nil {
		t.Fatalf("MarshalBinary failed: %v", err)
	}

	if string(data) != string(expected) {
		t.Errorf("MarshalBinary() = %v; want %v", data, expected)
	}
}

func TestUnmarshalBinary(t *testing.T) {
	var u UUID
	input := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	expected := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}

	err := u.UnmarshalBinary(input)
	if err != nil {
		t.Fatalf("UnmarshalBinary failed: %v", err)
	}

	if u != expected {
		t.Errorf("UnmarshalBinary() = %v; want %v", u, expected)
	}
}

func TestUnmarshalBinaryInvalidLength(t *testing.T) {
	var u UUID
	input := []byte{0x01, 0x23, 0x45} // Invalid length

	err := u.UnmarshalBinary(input)
	if err == nil {
		t.Fatalf("UnmarshalBinary should have failed for invalid length")
	}
}

func TestJSONSerialization(t *testing.T) {
	u := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}
	expected := `"01234567-89ab-cdef-fedc-ba9876543210"`

	jsonData, err := json.Marshal(u)
	if err != nil {
		t.Fatalf("JSON serialization failed: %v", err)
	}

	if string(jsonData) != expected {
		t.Errorf("JSON serialization = %s; want %s", jsonData, expected)
	}
}

func TestJSONDeserialization(t *testing.T) {
	var u UUID
	input := `"01234567-89ab-cdef-fedc-ba9876543210"`
	expected := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10}

	err := json.Unmarshal([]byte(input), &u)
	if err != nil {
		t.Fatalf("JSON deserialization failed: %v", err)
	}

	if u != expected {
		t.Errorf("JSON deserialization = %v; want %v", u, expected)
	}
}
