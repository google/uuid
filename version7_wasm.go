// Copyright 2026 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build js wasip1

package uuid

// hasSubMilliClock reports whether the platform's wall clock is guaranteed to
// have sub-millisecond resolution. On the wasm targets (js and wasip1) the
// resolution is host-dependent and sub-millisecond precision is not
// guaranteed, so v7 fills the rand_a field with random bits instead of a
// timestamp fraction.
const hasSubMilliClock = false
