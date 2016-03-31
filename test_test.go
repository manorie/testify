package testify

import (
	"testing"
)

func TestAddAuthor(t *testing.T) {
	test, _ := CreateNewTest()
	added, _ := test.AddAuthor("Alfonso")

	t.Errorf("%v", added)
}
