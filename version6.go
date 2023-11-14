// Copyright 2018 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "encoding/binary"

// NewV6 returns a Version 6 UUID based on the current NodeID and clock
// sequence, and the current time.  If the NodeID has not been set by SetNodeID
// or SetNodeInterface then it will be set automatically.  If the NodeID cannot
// be set NewV6 returns nil.  If clock sequence has not been set by
// SetClockSequence then it will be set automatically.  If GetTime fails to
// return the current NewV6 returns nil and an error.
func NewV6() (UUID, error) {
	var uuid UUID
	now, seq, err := GetTime()
	if err != nil {
		return uuid, err
	}

	binary.BigEndian.PutUint64(uuid[0:], uint64(now))
	binary.BigEndian.PutUint16(uuid[8:], seq)

	uuid[6] = 0x60 | (uuid[6] & 0x0F)
	uuid[8] = 0x80 | (uuid[8] & 0x3F)

	nodeMu.Lock()
	if nodeID == zeroID {
		setNodeInterface("")
	}
	copy(uuid[10:], nodeID[:])
	nodeMu.Unlock()

	return uuid, nil
}
