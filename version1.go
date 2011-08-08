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
)

// NewUUID returns a Version 1 UUID based on the current NodeID and clock
// sequence, and the current time.  If the NodeID has not been set by SetNodeID
// or SetNodeInterface then it will be set automatically.  If the NodeID cannot
// be set NewUUID returns nil.  If clock sequence has not been set by
// SetClockSequence then it will be set automatically.  If GetTime fails to
// return the current NewUUID returns nil.
func NewUUID() UUID {
	if nodeID == nil {
		SetNodeInterface("")
	}

	now, err := GetTime()
	if err != nil {
		return nil
	}

	uuid := make([]byte, 16)

	time_low := uint32(now & 0xffffffff)
	time_mid := uint16((now >> 32) & 0xffff)
	time_hi := uint16((now >> 48) & 0x0fff)
	time_hi |= 0x1000 // Version 1

	binary.BigEndian.PutUint32(uuid[0:], time_low)
	binary.BigEndian.PutUint16(uuid[4:], time_mid)
	binary.BigEndian.PutUint16(uuid[6:], time_hi)
	binary.BigEndian.PutUint16(uuid[8:], clock_seq)
	copy(uuid[10:], nodeID)

	return uuid
}
