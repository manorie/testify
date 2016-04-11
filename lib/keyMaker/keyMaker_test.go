package keyMaker

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tokens := make([]string, 1000, 1000)

	for i, _ := range tokens {
		tokens[i] = Generate()
	}
	for i, t0 := range tokens {
		for y, t1 := range tokens {
			if i != y && t0 == t1 {
				t.Fatalf("Token t1 : %v is exactly same as t0 : %v", t1, t0)
			}
		}
	}
}
