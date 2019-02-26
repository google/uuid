package uuid

import (
	"math/rand"
	"testing"
	"time"
)

func TestUuidSources(t *testing.T) {
	currentTime := time.Now().UnixNano()
	uuidSourceA := NewSource(rand.New(rand.NewSource(currentTime)))
	uuidSourceB := NewSource(rand.New(rand.NewSource(currentTime)))
	
	for i := 0; i < 10; i++ {
		if uuidSourceA.New().String() != uuidSourceB.New().String() {
			t.Error("Uuid values are not reproducaible!")
		}
	}
	
	uuidSourceA = NewSource(rand.New(rand.NewSource(123)))
	uuidSourceB = NewSource(rand.New(rand.NewSource(456)))
	

	for i := 0; i < 10; i++ {
		if uuidSourceA.New().String() == uuidSourceB.New().String() {
			t.Error("Uuid values should not match!")
		}
	}

}