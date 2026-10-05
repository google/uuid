// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.27

package uuid

// NewV4 returns a Version 4 UUID.
//
// NewV4 is equivalent to NewRandom. On Go 1.27 and later, NewV4
// delegates to the standard library's uuid.NewV4 implementation.
func NewV4() (UUID, error) {
	return NewRandom()
}
