package uuid

import (
	"io"
	"crypto/rand"
)

// A UuidSource holds a random number generator and generates UUIDs using it as its source.
//
// It is useful when a process need its own random number generator, 
// e.g. when running some processes concurrently.
type UuidSource struct {
	rander io.Reader
}

// Creates a new UuidSource which holds its own random number generator.
//
// Calling NewSource with nil sets the random number generator to a default
// generator.
func NewSource(r io.Reader) UuidSource {
	var uuidSource UuidSource
	uuidSource.SetRand(r)
	return uuidSource

}

// SetRand sets the random number generator of the UuidSource to r, which implements io.Reader.
// If r.Read returns an error when the package requests random data then
// a panic will be issued.
//
// Calling SetRand with nil sets the random number generator to a default
// generator.
func (uuidSource *UuidSource) SetRand(r io.Reader) {
	if r == nil {
		uuidSource.rander = rand.Reader
		return
	}
	uuidSource.rander = r
}

// NewRandom returns a Random (Version 4) UUID based on the random number generator in the UuidSource.
//
// See more detailed explanation here: https://godoc.org/github.com/google/uuid#NewRandom
func (uuidSource UuidSource) NewRandom() (UUID, error) {
	return newRandom(uuidSource.rander)
}

// New creates a new random UUID  based on the random number generator in the UuidSource or panics.  New is equivalent to
// the expression
//
//    uuid.Must(uuid.NewRandom())
func (uuidSource UuidSource) New() UUID {
	return Must(uuidSource.NewRandom())
}
