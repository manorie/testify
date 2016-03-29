package testify

import (
	"testing"
)

func TestCreateEmptyTest(t *testing.T) {
	test, _ := ConstructASafePath()

	t.Errorf("%v", test)
}
