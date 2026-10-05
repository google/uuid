// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.27

package uuid

// NewV7 returns a Version 7 UUID based on the current time (Unix Epoch).
// Uses the randomness pool if it was enabled with EnableRandPool.
// On error, NewV7 returns Nil and an error.
func NewV7() (UUID, error) {
	uuid, err := NewRandom()
	if err != nil {
		return uuid, err
	}
	makeV7(uuid[:])
	return uuid, nil
}
