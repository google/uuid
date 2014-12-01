// Copyright 2014 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"encoding/json"
	"testing"
)

var testUUID = Parse("f47ac10b-58cc-0372-8567-0e02b2c3d479")

func TestJSON(t *testing.T) {
	type S struct {
		ID UUID
	}
	s1 := S{testUUID}
	data, err := json.Marshal(&s1)
	if err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := json.Unmarshal(data, &s2); err != nil {
		t.Fatal(err)
	}
	if !Equal(s1.ID, s2.ID) {
		t.Errorf("got UUID %v, want %v", s2.ID, s1.ID)
	}
}
