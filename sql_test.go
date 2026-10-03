// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"errors"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	badTypeTest := 6
	invalidTest := "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	byteTest := make([]byte, 16)
	byteTestUUID := MustParse(stringTest)
	copy(byteTest, byteTestUUID[:])

	// sunny day tests

	var uuid UUID
	err := (&uuid).Scan(stringTest)
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan([]byte(stringTest))
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan(byteTest)
	if err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err = (&uuid).Scan(badTypeTest)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}
	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = (&uuid).Scan(invalidTest)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid UUID returned an incorrect error message")
	}

	err = (&uuid).Scan(byteTest[:len(byteTest)-2])
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid byte UUID returned an incorrect error message")
	}

	// empty tests

	uuid = UUID{}
	var emptySlice []byte
	err = (&uuid).Scan(emptySlice)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	var emptyString string
	err = (&uuid).Scan(emptyString)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	uuid = UUID{}
	err = (&uuid).Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning nil")
		}
	}
}

func TestValue(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	uuid := MustParse(stringTest)
	val, _ := uuid.Value()
	if val != stringTest {
		t.Error("Value() did not return expected string")
	}
}

// TestScanErrorWrapping checks that Scan reports the underlying parse failure
// rather than flattening it into an opaque string. The package exports
// ErrInvalidLength, ErrInvalidUUIDFormat and ErrInvalidURNPrefix together with
// the IsInvalidLengthError helper precisely so that callers can classify parse
// errors with errors.Is/errors.As; Scan must not defeat that.
func TestScanErrorWrapping(t *testing.T) {
	testCases := []struct {
		name  string
		text  string
		src   interface{}
		match func(error) bool
	}{
		{"string/invalid-length", "12345", "12345", IsInvalidLengthError},
		{"bytes/invalid-length", "12345", []byte("12345"), IsInvalidLengthError},
		{"string/invalid-format", "12345678gabc1234abcd1234abcd1234", "12345678gabc1234abcd1234abcd1234", func(err error) bool {
			return errors.Is(err, ErrInvalidUUIDFormat)
		}},
		{"bytes/invalid-format", "12345678gabc1234abcd1234abcd1234", []byte("12345678gabc1234abcd1234abcd1234"), func(err error) bool {
			return errors.Is(err, ErrInvalidUUIDFormat)
		}},
		{"string/invalid-urn-prefix", "urn:test:123e4567-e89b-12d3-a456-426655440000", "urn:test:123e4567-e89b-12d3-a456-426655440000", func(err error) bool {
			return errors.Is(err, ErrInvalidURNPrefix)
		}},
		{"bytes/invalid-urn-prefix", "urn:test:123e4567-e89b-12d3-a456-426655440000", []byte("urn:test:123e4567-e89b-12d3-a456-426655440000"), func(err error) bool {
			return errors.Is(err, ErrInvalidURNPrefix)
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var uuid UUID
			err := uuid.Scan(tc.src)
			if err == nil {
				t.Fatalf("Scan(%v) succeeded, want error", tc.src)
			}
			if !tc.match(err) {
				t.Errorf("Scan(%v) = %v: underlying parse error type was lost", tc.src, err)
			}
			// The error message itself must be unchanged.
			_, parseErr := Parse(tc.text)
			if want := "Scan: " + parseErr.Error(); err.Error() != want {
				t.Errorf("Scan(%v) message = %q, want %q", tc.src, err.Error(), want)
			}
		})
	}
}
