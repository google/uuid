// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "testing"

func TestNewV4(t *testing.T) {
	SetRand(nil)
	m := make(map[string]bool)
	for x := 1; x < 128; x++ {
		uuid, err := NewV4()
		if err != nil {
			t.Fatalf("could not create UUID: %v", err)
		}
		s := uuid.String()
		if m[s] {
			t.Errorf("NewV4 returned duplicated UUID %s", s)
		}
		m[s] = true
		if v := uuid.Version(); v != 4 {
			t.Errorf("UUID of version %s", v)
		}
		if uuid.Variant() != RFC4122 {
			t.Errorf("UUID is variant %d", uuid.Variant())
		}
	}
}
