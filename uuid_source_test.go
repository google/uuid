package uuid

import (
	"testing"
	"strings"
	
	regen "github.com/zach-klippenstein/goregen"
)

func TestUuidSources(t *testing.T) {
	
	myString, _ := regen.Generate("[a-zA-Z]{1000}")

	// Two identical sources, should give same sequence
	uuidSourceA := NewSource(strings.NewReader(myString))
	uuidSourceB := NewSource(strings.NewReader(myString))
	
	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 != uuid2 {
			t.Errorf("expected duplicates, got %q and %q", uuid1, uuid2)
		}
	}
	
	// Set rander with nil, each source will be random
	uuidSourceA.SetRand(nil)
	uuidSourceB.SetRand(nil)

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q", uuid1)
		}
	}
	
	// Set rander to rand source with same seed, should give same sequence
	uuidSourceA.SetRand(strings.NewReader(myString))
	uuidSourceB.SetRand(strings.NewReader(myString))

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 != uuid2 {
			t.Errorf("expected duplicates, got %q and %q", uuid1, uuid2)
		}
	}

	// Set rander to rand source with different seeds, should not give same sequence
	uuidSourceA.SetRand(strings.NewReader("456" + myString))
	uuidSourceB.SetRand(strings.NewReader("myString" + myString))

	for i := 0; i < 10; i++ {
		uuid1 := uuidSourceA.New()
		uuid2 := uuidSourceB.New()
		if uuid1 == uuid2 {
		t.Errorf("unexpected duplicates, got %q", uuid1)
		}
	}

}