// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.27

package uuid

import (
	"testing"
	stduuid "uuid"
)

func TestGo127ShimsConvertToStdlib(t *testing.T) {
	v4, err := NewV4()
	if err != nil {
		t.Fatalf("NewV4 returned error: %v", err)
	}
	if got := UUID(stduuid.UUID(v4)); got != v4 {
		t.Errorf("NewV4 did not survive conversion to and from stdlib uuid: got %s want %s", got, v4)
	}
	if v4.Version() != 4 {
		t.Errorf("NewV4 version = %s, want 4", v4.Version())
	}

	v7, err := NewV7()
	if err != nil {
		t.Fatalf("NewV7 returned error: %v", err)
	}
	if got := UUID(stduuid.UUID(v7)); got != v7 {
		t.Errorf("NewV7 did not survive conversion to and from stdlib uuid: got %s want %s", got, v7)
	}
	if v7.Version() != 7 {
		t.Errorf("NewV7 version = %s, want 7", v7.Version())
	}
}
