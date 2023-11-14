// Copyright 2018 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"encoding/binary"
	"io"
)

func unixepoch() (uint64, error) {
	now, _, err := GetTime()
	if err != nil {
		return 0, err
	}

	sec, nsec := now.UnixTime()
	return uint64(sec*1000 + nsec/1000000), nil
}

func NewV7FromReader(r io.Reader) (UUID, error) {
	var uuid UUID

	t, err := unixepoch()
	if err != nil {
		return uuid, err
	}

	binary.BigEndian.PutUint64(uuid[0:], t<<16)
	_, err = io.ReadFull(r, uuid[6:])
	if err != nil {
		return Nil, err
	}

	uuid[6] = 0x70 | (uuid[6] & 0x0F)
	uuid[8] = 0x80 | (uuid[8] & 0x3F)

	return uuid, nil
}

func newV7FromPool() (UUID, error) {
	var uuid UUID

	t, err := unixepoch()
	if err != nil {
		return uuid, err
	}

	if err = fill16BytesFromPool(uuid[:]); err != nil {
		return Nil, err
	}

	binary.BigEndian.PutUint32(uuid[0:], uint32(t>>16))
	binary.BigEndian.PutUint16(uuid[4:], uint16(t))

	uuid[6] = 0x70 | (uuid[6] & 0x0F)
	uuid[8] = 0x80 | (uuid[8] & 0x3F)

	return uuid, nil
}

// NewV7 returns a Version 7 UUID based on the current time.
// If GetTime fails to  return the current NewV7 returns nil and an error.
func NewV7() (UUID, error) {
	if !poolEnabled {
		return NewV7FromReader(rander)
	}
	return newV7FromPool()
}
