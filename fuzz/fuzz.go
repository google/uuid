// Copyright 2020 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuzz

import "github.com/google/uuid"

func Fuzz(data []byte) int {
	_, err := uuid.Parse(string(data))
	if err != nil {
		return 0
	}
	return 1
}
