package uuid

import (
	"io"
	"crypto/rand"
)

type UuidSource struct {
	rander io.Reader
}

func NewSource(r io.Reader) UuidSource {
	var uuidSource UuidSource
	uuidSource.SetRand(r)
	return uuidSource

}

func (uuidSource *UuidSource) SetRand(r io.Reader) {
	if r == nil {
		uuidSource.rander = rand.Reader
		return
	}
	uuidSource.rander = r
}

func (uuidSource UuidSource) NewRandom() (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(uuidSource.rander, uuid[:])
	if err != nil {
		return Nil, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

func (uuidSource UuidSource) New() UUID {
	return Must(uuidSource.NewRandom())
}
