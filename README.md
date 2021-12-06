# uuid ![build status](https://travis-ci.org/google/uuid.svg?branch=master)
The uuid package generates and inspects UUIDs based on
[RFC 4122](http://tools.ietf.org/html/rfc4122)
and DCE 1.1: Authentication and Security Services.

This package is based on the github.com/pborman/uuid package (previously named
code.google.com/p/go-uuid).  It differs from these earlier packages in that
a UUID is a 16 byte array rather than a byte slice.  One loss due to this
change is the ability to represent an invalid UUID (vs a NIL UUID).

###### Install
`go get github.com/google/uuid`

###### Example usage

```
package main

import (
  "fmt"

  "github.com/google/uuid"
)

func main() {

  // uuid in uuid.UUID form
  uid := uuid.New()
  fmt.Println(uid)

  // uuid v4 in a string form
  uidStr := uuid.NewString()
  fmt.Println(uidStr)

  // urn version urn:uuid:xxx...
  urnUuidStr := uuid.New().URN()
  fmt.Println(urnUuidStr)

  // Parse uuid
  parsedUuid, err := uuid.Parse("138e62c7-1f88-42cb-8185-8d9e07918e84")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(parsedUuid)
  }

}

```


###### Documentation
[![GoDoc](https://godoc.org/github.com/google/uuid?status.svg)](http://godoc.org/github.com/google/uuid)

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here:
http://pkg.go.dev/github.com/google/uuid
