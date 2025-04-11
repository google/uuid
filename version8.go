// Copyright 2024 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"encoding/binary"
	"io"
	"time"
)

var (
	lastTimestamp uint64
	sequence      uint16
)

// NewV8 generates a version 8 UUID.
// It sets the version and variant fields and fills the rest with random data.
// It allows embedding of user-defined data while maintaining the UUID structure.
//
// The layout includes:
// - custom_a: 48 bits (user-defined)
// - ver: 4 bits (set to 0b1000 for version 8)
// - custom_b: 12 bits (user-defined)
// - var: 2 bits (set to 0b10 for RFC 4122 variant)
// - custom_c: 62 bits (user-defined)
//
// For details, see https://datatracker.ietf.org/doc/html/rfc9562#name-uuid-version-8
//
// NewV8 generates a UUID version 8 with completely random user-defined fields.
// It uses the randomness pool if enabled or falls back to a secure random source.
// On error, NewV8 returns Nil and an error.
func NewV8() (UUID, error) {
	uuid, err := NewRandom()
	if err != nil {
		return uuid, err
	}
	uuid[6] = (uuid[6] & 0x0F) | 0x80 // Set version to 8
	uuid[8] = (uuid[8] & 0x3F) | 0x80 // Set variant to RFC 4122
	return uuid, nil
}

// NewV8FromReader generates a version 8 UUID with user-defined custom_a and custom_b.
// It uses random bits for custom_c if no random reader is provided.
// On error, NewV8FromReader returns Nil and an error.
func NewV8FromReader(customA, customB uint64, random io.Reader) (UUID, error) {
	var uuid UUID

	// Encode custom_a (48 bits)
	binary.BigEndian.PutUint64(uuid[:8], customA)
	copy(uuid[:6], uuid[2:8]) // Retain only the lower 48 bits

	// Set version and custom_b (12 bits)
	uuid[6] = (uuid[6] & 0x0F) | 0x80
	binary.BigEndian.PutUint16(uuid[6:8], uint16(customB)|0x8000)

	// Fill custom_c (62 bits)
	if random == nil {
		random = rander
	}
	if _, err := io.ReadFull(random, uuid[8:]); err != nil {
		return Nil, err
	}
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return uuid, nil
}

// NewV8TimeBased generates a version 8 UUID with a time-based custom_a field.
// It ensures uniqueness using a sequence number for UUIDs created in the same nanosecond.
func NewV8TimeBased(random io.Reader) (UUID, error) {
	var uuid UUID
	timestamp := uint64(time.Now().UnixNano())

	timeMu.Lock()

	if timestamp == lastTimestamp {
		sequence++
	} else {
		lastTimestamp = timestamp
		sequence = 0
	}

	timeMu.Unlock()

	// Encode timestamp into custom_a (48 bits)
	binary.BigEndian.PutUint64(uuid[:8], timestamp)
	copy(uuid[:6], uuid[2:8])

	// Set version and variant
	uuid[6] = (uuid[6] & 0x0F) | 0x80
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	// Add sequence to custom_c (16 bits)
	binary.BigEndian.PutUint16(uuid[8:], sequence)

	// Fill the rest with custom_c
	if random == nil {
		for i := 10; i < 16; i++ {
			uuid[i] = 0
		}
	} else if _, err := io.ReadFull(random, uuid[10:]); err != nil {
		return Nil, err
	}

	return uuid, nil
}
