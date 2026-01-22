// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"encoding/binary"
	"io"
	chacha8RandV2 "math/rand/v2"
)

// New creates a new random UUID or panics.  New is equivalent to
// the expression
//
//	uuid.Must(uuid.NewRandom())
func New() UUID {
	return Must(NewRandom())
}

// NewString creates a new random UUID and returns it as a string or panics.
// NewString is equivalent to the expression
//
//	uuid.New().String()
func NewString() string {
	return Must(NewRandom()).String()
}

// NewRandom returns a Random (Version 4) UUID.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// Uses the randomness pool if it was enabled with EnableRandPool.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//	Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//	hit by a meteorite is estimated to be one chance in 17 billion, that
//	means the probability is about 0.00000000006 (6 × 10−11),
//	equivalent to the odds of creating a few tens of trillions of UUIDs in a
//	year and having one duplicate.
func NewRandom() (UUID, error) {
	if _, ok := rander.(defaultRandReader); ok {
		return fastRandV4DefaultZeroAlloc(), nil
	}

	if !poolEnabled {
		return NewRandomFromReader(rander)
	}
	return newRandomFromPool()
}

// NewRandomFromReader returns a UUID based on bytes read from a given io.Reader.
func NewRandomFromReader(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return Nil, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

func newRandomFromPool() (UUID, error) {
	var uuid UUID
	poolMu.Lock()
	if poolPos == randPoolSize {
		_, err := io.ReadFull(rander, pool[:])
		if err != nil {
			poolMu.Unlock()
			return Nil, err
		}
		poolPos = 0
	}
	copy(uuid[:], pool[poolPos:(poolPos+16)])
	poolPos += 16
	poolMu.Unlock()

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

// Singleton with no state that implements
// io.Reader and uses the default rand/v2
// package to read bytes.
type defaultRandReader struct{}

// Read fills the provided byte slice `b` with random data using a
// cryptographically secure random number generator.
func (d defaultRandReader) Read(b []byte) (n int, err error) {
	var num uint64
	numByteIndex := 0

	for i := range b {
		if numByteIndex == 0 {
			num = chacha8RandV2.Uint64()
		}
		b[i] = byte(num >> (numByteIndex * 8))
		numByteIndex = (numByteIndex + 1) % 8
	}

	return len(b), nil
}

func fastRandV4DefaultZeroAlloc() UUID {
	var uuid UUID
	hi, low := chacha8RandV2.Uint64(), chacha8RandV2.Uint64()
	binary.BigEndian.PutUint64(uuid[0:8], hi)
	binary.BigEndian.PutUint64(uuid[8:16], low)

	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}
