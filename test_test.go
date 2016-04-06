package testify

import (
	"testing"
)

func TestAddAuthor(t *testing.T) {
	test, _ := CreateNewTest()
	test.AddAuthor("Alfonso")
	test.AddTitle("Back to the future")
	test.AddAuthorEmail("mcetin.cm@gmail.com")
	test.SetTimeLimit(20)
	test.SetAnswerSize(4)
	test.SetExpireInDays(10)
	//t.Errorf("%v", test)
}
