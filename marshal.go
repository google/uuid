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

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	copy(uuid[:], data)
	return nil
}

// MarshalBinaryLittleEndian implements encoding.BinaryMarshaler.
func (uuid UUID) MarshalBinaryLittleEndian() ([]byte, error) {
	var uuidLittleEndian UUID = UUID {
		// convert first 3 fields from big-endian to little-endian
		uuid[3], uuid[2], uuid[1], uuid[0],
		uuid[5], uuid[4],
		uuid[7], uuid[6],
	}
	// all the rest fields keep byte order
	copy(uuidLittleEndian[8:], uuid[8:])
	return uuidLittleEndian[:], nil
}

// UnmarshalBinaryLittleEndian implements encoding.BinaryUnmarshaler.
func (uuid *UUID) UnmarshalBinaryLittleEndian(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	// convert first 3 fields from little-endian to big-endian
	uuid[0] = data[3]; uuid[1] = data[2]; uuid[2] = data[1]; uuid[3] = data[0]
	uuid[4] = data[5]; uuid[5] = data[4]
	uuid[6] = data[7]; uuid[7] = data[6]
	// all the rest fields keep byte order
	copy(uuid[8:], data[8:])
	return nil
}
