// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uuid generates and inspects UUIDs.
//
// UUIDs are based on RFC 4122 and DCE 1.1: Authentication and Security
// Services.
//
// A UUID is a 16 byte (128 bit) array.  UUIDs may be used as keys to
// maps or compared directly.
// 
// Sample usage:
// 
//       package main
//
//       import (
//         "fmt"
//
//         "github.com/google/uuid"
//       )
//
//       func main() {
//         id := uuid.New()
//         fmt.Println("id:", id.String())
//       }
package uuid
