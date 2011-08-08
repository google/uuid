// Copyright 2011 Google Inc.  All rights reserved.
// Author: borman@google.com (Paul Borman)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The uuid package generates and inspects UUIDs based on RFC 4122 and
// DCE 1.1: Authentication and Security Services.
package uuid

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
)

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
type UUID []byte

// A Version represents a UUIDs version.
type Version byte

// A Variant represents a UUIDs variant.
type Variant byte

// Constants returned by Variant.
const (
	INVALID   = iota // Invalid UUID
	RFC4122          // The variant specified in RFC4122
	RESERVED         // Reserved, NCS backward compatibility.
	MICROSOFT        // Reserved, Microsoft Corporation backward compatibility.
	FUTURE           // Reserved for future definition.
)

var rander = rand.Reader // random function

// New returns a new random (version 4) UUID as a string.  It is a convenience
// function for NewRandom().String().
func New() string {
	return NewRandom().String()
}

// Decode decodes s into a UUID or returns nil.  Both the UUID form of
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded.
func Decode(s string) UUID {
	if len(s) == 36+9 {
		if strings.ToLower(s[:9]) != "urn:uuid:" {
			return nil
		}
		s = s[9:]
	} else if len(s) != 36 {
		return nil
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return nil
	}
	uuid := make([]byte, 16)
	for i, x := range []int{
		0, 2, 4, 6,
		9, 11,
		14, 16,
		19, 21,
		24, 26, 28, 30, 32, 34} {
		if v, ok := xtob(s[x:]); !ok {
			return nil
		} else {
			uuid[i] = v
		}
	}
	return uuid
}

// Equal returns true if uuid1 and uuid2 are equal.
func Equal(uuid1, uuid2 UUID) bool {
	return bytes.Equal(uuid1, uuid2)
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (uuid UUID) String() string {
	if uuid == nil || len(uuid) != 16 {
		return ""
	}
	b := []byte(uuid)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// URN returns the RFC 2141 URN form of uuid,
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx,  or "" if uuid is invalid.
func (uuid UUID) URN() string {
	if uuid == nil || len(uuid) != 16 {
		return ""
	}
	b := []byte(uuid)
	return fmt.Sprintf("urn:uuid:%08x-%04x-%04x-%04x-%012x",
		b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// Variant returns the variant encoded in uuid.  It returns INVALID if
// uuid is invalid.
func (uuid UUID) Variant() Variant {
	if len(uuid) != 16 {
		return INVALID
	}
	switch {
	case (uuid[8] & 0xc0) == 0x80:
		return RFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return MICROSOFT
	case (uuid[8] & 0xe0) == 0xe0:
		return FUTURE
	default:
		return RESERVED
	}
	panic("unreachable")
}

// Version returns the verison of uuid.  It returns false if uuid is not
// valid.
func (uuid UUID) Version() (Version, bool) {
	if len(uuid) != 16 {
		return 0, false
	}
	return Version(uuid[6] >> 4), true
}

func (v Version) String() string {
	if v > 15 {
		return fmt.Sprintf("BAD_VERSION_%d", v)
	}
	return fmt.Sprintf("VERSION_%d", v)
}

func (v Variant) String() string {
	switch v {
	case RFC4122:
		return "RFC4122"
	case RESERVED:
		return "RESERVED"
	case MICROSOFT:
		return "MICROSOFT"
	case FUTURE:
		return "FUTURE"
	case INVALID:
		return "INVALID"
	}
	return fmt.Sprintf("BAD_VARIANT_%d", v)
}

// SetRand sets the random number generator to r, which implents io.Reader.
// If r.Read returns an error when the package requests random data then
// a panic will be issued.
//
// Calling SetRand with nil sets the random number generator to the default
// generator.
func SetRand(r io.Reader) {
	if r == nil {
		rander = rand.Reader
		return
	}
	rander = r
}
