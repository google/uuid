// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"hash"
)

// Well known namespace IDs and UUIDs
var (
	NameSpaceDNS  = Must(Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceURL  = Must(Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceOID  = Must(Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
	NameSpaceX500 = Must(Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
	Nil           UUID // empty UUID, all zeros

	// The Max UUID is special form of UUID that is specified to have all 128 bits set to 1.
	MAX = UUID{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	}
)

// NewHash returns a new UUID derived from the hash of space concatenated with
// data generated by h.  The hash should be at least 16 byte in length.  The
// first 16 bytes of the hash are used to form the UUID.  The version of the
// UUID will be the lower 4 bits of version.  NewHash is used to implement
// NewMD5 and NewSHA1.
func NewHash(h hash.Hash, space UUID, data []byte, version int) UUID {
	h.Reset()
	h.Write(space[:]) //nolint:errcheck
	h.Write(data)     //nolint:errcheck
	s := h.Sum(nil)
	var uuid UUID
	copy(uuid[:], s)
	uuid[6] = (uuid[6] & 0x0f) | uint8((version&0xf)<<4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // RFC 4122 variant
	return uuid
}

// NewMD5 returns a new MD5 (Version 3) UUID based on the
// supplied name space and data.  It is the same as calling:
//
//	NewHash(md5.New(), space, data, 3)
func NewMD5(space UUID, data []byte) UUID {
	return NewHash(md5.New(), space, data, 3)
}

// NewSHA1 returns a new SHA1 (Version 5) UUID based on the
// supplied name space and data.  It is the same as calling:
//
//	NewHash(sha1.New(), space, data, 5)
func NewSHA1(space UUID, data []byte) UUID {
	return NewHash(sha1.New(), space, data, 5)
}
