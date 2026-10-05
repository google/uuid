// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.27

package uuid

import stduuid "uuid"

// NewV7 returns a Version 7 UUID based on the current time (Unix Epoch).
//
// On Go 1.27 and later, NewV7 delegates to the standard library's
// uuid.NewV7 implementation. On error, NewV7 returns Nil and an error;
// the standard implementation does not report an error, so the error
// returned by this shim is always nil.
//
//go:fix inline
func NewV7() (UUID, error) {
	return UUID(stduuid.NewV7()), nil
}
