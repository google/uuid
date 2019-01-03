// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var (
	// By default, new (recommended) UUID subtype/kind (0x04) is used, see
	// https://studio3t.com/knowledge-base/articles/mongodb-best-practices-uuid-data/#binary-subtypes-0x03-and-0x04 and
	// http://bsonspec.org/spec.html for details. Can be changed with SetBSONKind.
	bsonKind byte = 0x04
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
	if err == nil {
		*uuid = id
	}
	return err
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

// GetBSON implements bson.Getter for marshaling UUID in BSON binary UUID format.
// By default Kind of 0x04 (new UUID) will be used. Can be changed with SetBSONKind.
func (uuid UUID) GetBSON() (interface{}, error) {
	toMarshal := bson.Binary{
		Kind: bsonKind,
		Data: uuid[:],
	}

	return toMarshal, nil
}

// SetBSON implements bson.Setter for unmarshaling UUID from BSON binary UUID format.
// By default Kind of 0x04 (new UUID) will be used. Can be changed with SetBSONKind.
func (uuid *UUID) SetBSON(raw bson.Raw) error {
	var toUnmarshal bson.Binary

	err := raw.Unmarshal(&toUnmarshal)
	if err != nil {
		return err
	}

	*uuid, err = FromBytes(toUnmarshal.Data)

	return err
}

// SetBSONKind changes BSON UUID Kind which will be used by GetBSON and SetBSON. Only values of
// 0x03 (Legacy UUID) or 0x04 (new UUID) can be used, SetBSONKind returns error when requested kind is different.
func SetBSONKind(kind byte) error {
	if kind < 0x03 || kind > 0x04 {
		return errors.New("requested BSON UUID kind is not allowed")
	}

	bsonKind = kind

	return nil
}
