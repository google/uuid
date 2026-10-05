// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.27

package uuid

import stduuid "uuid"

// NewV4 returns a Version 4 UUID.
//
// On Go 1.27 and later, NewV4 delegates to the standard library's
// uuid.NewV4 implementation.
//
//go:fix inline
func NewV4() (UUID, error) {
	return UUID(stduuid.NewV4()), nil
}
