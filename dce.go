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

package uuid

import (
	"encoding/binary"
	"fmt"
	"os"
)

// A Domain represents a Version 2 domain
type Domain byte

// Domain constants for DCE Security (Version 2) UUIDs.
const (
	DOMAIN_PERSON = 0
	DOMAIN_GROUP  = 1
	DOMAIN_ORG    = 2
)

// NewDCESecurity returns a DCE Security (Version 2) UUID.
//
// The domain should be one of DOMAIN_PERSON, DOMAIN_GROUP or DOMAIN_ORG.
// On a POSIX system the id should be the users UID for the DOMAIN_PERSON
// domain and the users GID for the DOMAIN_GROUP.  The meaning of id for
// the domain DOMAIN_ORG or on non-POSIX systems is site defined.
//
// For a given domain/id pair the same token may be returned for up to
// 7 minutes and 10 seconds.
func NewDCESecurity(domain Domain, id uint32) UUID {
	uuid := NewUUID()
	if uuid != nil {
		uuid[6] = (uuid[6] & 0x0f) | 0x20 // Version 2
		uuid[9] = byte(domain)
		binary.BigEndian.PutUint32(uuid[0:], id)
	}
	return uuid
}

// NewDCEGroup returns a DCE Security (Version 2) UUID in the group
// domain with the id returned by os.Getuid.
//
//  NewDCESecurity(DOMAIN_PERSON, uint32(os.Getuid()))
func NewDCEPerson() UUID {
	return NewDCESecurity(DOMAIN_PERSON, uint32(os.Getuid()))
}

// NewDCEPerson returns a DCE Security (Version 2) UUID in the group
// domain with the id returned by os.Getgid.
//
//  NewDCESecurity(DOMAIN_GROUP, uint32(os.Getgid()))
func NewDCEGroup() UUID {
	return NewDCESecurity(DOMAIN_GROUP, uint32(os.Getgid()))
}

// Domain returns the domain for a Version 2 UUID or false.
func (uuid UUID) Domain() (Domain, bool) {
	if v, _ := uuid.Version(); v != 2 {
		return 0, false
	}
	return Domain(uuid[9]), true
}

// Id returns the id for a Version 2 UUID or false.
func (uuid UUID) Id() (uint32, bool) {
	if v, _ := uuid.Version(); v != 2 {
		return 0, false
	}
	return binary.BigEndian.Uint32(uuid[0:4]), true
}

func (d Domain) String() string {
	switch d {
	case DOMAIN_PERSON:
		return "DOMAIN_PERSON"
	case DOMAIN_GROUP:
		return "DOMAIN_GROUP"
	case DOMAIN_ORG:
		return "DOMAIN_ORG"
	}
	return fmt.Sprintf("DOMAIN_%d", d)
}
