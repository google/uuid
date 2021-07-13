// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"fmt"
	"io"
)

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

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (uuid *UUID) UnmarshalGQL(v interface{}) (err error) {
	if s, ok := v.(string); ok {
		id, err := Parse(s)
		if err == nil {
			*uuid = id
		}
		return err
	}

	return fmt.Errorf("invalid UUID type, value must string")
}

// MarshalGQL implements the graphql.Marshaler interface
func (uuid UUID) MarshalGQL(w io.Writer) {
	var gql [38]byte
	gql[0] = '"'
	encodeHex(gql[1:37], uuid)
	gql[37] = '"'
	w.Write(gql[:])
}
