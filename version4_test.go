package uuid

import (
	"bytes"
	"testing"
)

func TestVersion4(t *testing.T) {
	v4 := &Version4{}

	m := make(map[UUID]bool)
	for x := 1; x < 32; x++ {
		s := v4.NewUUID()
		if m[s] {
			t.Errorf("NewUUID returned duplicated UUID %s", s)
		}
		m[s] = true
		uuid, err := Parse(s.String())
		if err != nil {
			t.Errorf("NewUUID.String() returned %q which does not decode", s)
			continue
		}
		if v := uuid.Version(); v != 4 {
			t.Errorf("Random UUID of version %s", v)
		}
		if uuid.Variant() != RFC4122 {
			t.Errorf("Random UUID is variant %d", uuid.Variant())
		}
	}
}

func TestVersion4Rand(t *testing.T) {
	v4 := &Version4{
		Rand: bytes.NewReader([]byte{
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
			0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
		}),
	}

	uuid := v4.NewString()
	if uuid != "00010203-0405-4607-8809-0a0b0c0d0e0f" {
		t.Errorf("NewString returned unexpected v4 UUID: %q", uuid)
	}
}
