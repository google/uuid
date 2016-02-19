// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

// New returns a Random (Version 4) UUID or panics.
//
// The strength of the UUIDs is based on the strength of the crypto/rand
// package.
//
// A note about uniqueness derived from from the UUID Wikipedia entry:
//
//  Randomly generated UUIDs have 122 random bits.  One's annual risk of being
//  hit by a meteorite is estimated to be one chance in 17 billion, that
//  means the probability is about 0.00000000006 (6 × 10−11),
//  equivalent to the odds of creating a few tens of trillions of UUIDs in a
//  year and having one duplicate.
func New() UUID {
	var uuid UUID
	randomBits([]byte(uuid[:]))
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid
}
