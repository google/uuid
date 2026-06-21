// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js,!wasip1

package uuid

// hasSubMilliClock reports whether the platform's wall clock is guaranteed to
// have sub-millisecond resolution. Non-wasm targets provide it, so v7 derives
// the rand_a field from a real sub-millisecond timestamp fraction.
const hasSubMilliClock = true
