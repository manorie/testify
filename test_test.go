package testify

import (
	"fmt"
	"testing"
)

func TestAddAuthor(t *testing.T) {
	db, err := OpenDb()
	if err != nil {
		t.Fatalf("%v", err)
	}

	test, _ := CreateNewTest()
	test.AddAuthor("Alfonso")
	test.AddTitle("Back to the future")
	test.AddAuthorEmail("mcetin.cm@gmail.com")
	test.SetTimeLimit(20)
	test.SetAnswerSize(4)
	test.SetExpireInDays(10)

	err = test.WriteTo(db)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(test)

	db.Close()
}
