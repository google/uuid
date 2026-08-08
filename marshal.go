// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "fmt"

// MarshalText implements encoding.TextMarshaler.
func (uuid UUID) MarshalText() ([]byte, error) {
	var js [36]byte
	encodeHex(js[:], uuid)
	return js[:], nil
}

// AppendText implements encoding.TextAppender.
// It appends the 36-byte textual form of uuid to b and returns the extended buffer.
func (uuid UUID) AppendText(b []byte) ([]byte, error) {
	// Grow by 36 in one append, then fill the new space in place.
	n := len(b)
	b = append(b, make([]byte, 36)...)
	encodeHex(b[n:], uuid)
	return b, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (uuid *UUID) UnmarshalText(data []byte) error {
	id, err := ParseBytes(data)
	if err != nil {
		return err
	}
	*uuid = id
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (uuid UUID) MarshalBinary() ([]byte, error) {
	return uuid[:], nil
}

// AppendBinary implements encoding.BinaryAppender.
// It appends the 16-byte binary form of uuid to b and returns the extended buffer.
func (uuid UUID) AppendBinary(b []byte) ([]byte, error) {
	return append(b, uuid[:]...), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	copy(uuid[:], data)
	return nil
}
