package yenc

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	b := make([]byte, 1024)
	rand.Read(b)
	fmt.Printf("encoding % x", b)
	encoded := Encode(b)
	decoded := Decode(encoded)
	if string(b) != string(decoded) {
		t.Errorf("Decode() = % x, want: % x", decoded, b)
	}
}
