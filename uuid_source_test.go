package uuid

import (
	"testing"
	"fmt"
	"strings"
	"math/rand"

)

func TestUuidSources(t *testing.T) {

	uuidSourceA := NewSource(rand.New(rand.NewSource(34576)))
	uuidSourceB := NewSource(rand.New(rand.NewSource(34576)))
	
	var uuidStrA, uuidStrB string
	fmt.Printf("Random values: ")
	for i := 0; i < 10; i++ {
		uuidStrA += uuidSourceA.New().String() + "."
	}
	fmt.Printf("%v\n", uuidStrA)
	
	fmt.Printf("Random values: ")
	for i := 0; i < 10; i++ {
		uuidStrB += uuidSourceB.New().String() + "."
	}
	fmt.Printf("%v\n", uuidStrB)
	
	if !strings.EqualFold(uuidStrA, uuidStrB) {
		t.Error("Uuid values are not reproducaible!")
	}
	
	uuidSourceA = NewSource(rand.New(rand.NewSource(66)))
	uuidSourceB = NewSource(rand.New(rand.NewSource(77)))
	

	uuidStrA = ""
	uuidStrB = ""
	fmt.Printf("Random values: ")
	for i := 0; i < 10; i++ {
		uuidStrA += uuidSourceA.New().String() + "."
	}
	fmt.Printf("%v\n", uuidStrA)
	
	fmt.Printf("Random values: ")
	for i := 0; i < 10; i++ {
		uuidStrB += uuidSourceB.New().String() + "."
	}
	fmt.Printf("%v\n", uuidStrB)
	
	if strings.EqualFold(uuidStrA, uuidStrB) {
		t.Error("Uuid values should not be reproducaible!")
	}
	
}