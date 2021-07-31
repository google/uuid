// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"crypto/rand"
	"io"
)

// New creates a new random UUID or panics.  New is equivalent to
// the expression
//
//    uuid.Must(uuid.NewRandom())
//
// Deprecated: Use *Version4.NewUUID() instead.
func New() UUID {
	return Must(NewRandom())
}

// NewString creates a new random UUID and returns it as a string or panics.
// NewString is equivalent to the expression
//
//    uuid.New().String()
//
// Deprecated: Use *Version4.NewString() instead.
func NewString() string {
	return Must(NewRandom()).String()
}

// NewRandom returns a Random (Version 4) UUID.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
//
// Deprecated: Use *Version4.New() instead.
func NewRandom() (UUID, error) {
	g := &Version4{Rand: rander}
	return g.New()
}

// NewRandomFromReader returns a UUID based on bytes read from a given io.Reader.
//
// Deprecated: Use *Version4.New() instead.
func NewRandomFromReader(r io.Reader) (UUID, error) {
	g := &Version4{Rand: r}
	return g.New()
}

// A Version4 generates Version 4 UUIDs.
type Version4 struct {
	// Rand provides the source of entropy for generation random UUIDs.
	// If Rand is nil, crypto/rand.Reader will be used.
	//
	// The thread-safety of the Version4 is contingent upon that of the Rand.
	//
	// Clients wishing to sacrifice security for performance should consider
	// using a bufio.Reader or similar.
	Rand io.Reader
}

// New generates a new a Version 4 UUID.
//
// The strength/randomness of the UUID is based on the strength of v4.Rand.
//
// A note about uniqueness derived from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
func (v4 *Version4) New() (UUID, error) {
	r := v4.Rand
	if r == nil {
		r = rand.Reader
	}
	return newV4(r)
}

// NewUUID generates a new Version 4 UUID, or panics.
// It is equivalent to:
//  uuid.Must(v4.New())
func (v4 *Version4) NewUUID() UUID {
	return Must(v4.New())
}

// NewString generates a new Version 4 UUID as a string, or panics.
// It is equivalent to:
//  v4.NewUUID().String()
func (v4 *Version4) NewString() string {
	return v4.NewUUID().String()
}

func newV4(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return Nil, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}
