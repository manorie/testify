package test

import (
	"testing"
)

func TestPackage(t *testing.T) {
	test := BuildEmptyTest()
	test.AddTitle("Meditations")
	test.AddAuthor("Marcus Aurelius")
	test.AddAuthorEmail("marcus@rome.it")
	test.SetTimeLimit(3)
	test.SetAnswerSize(4)
	test.SetExpireInDays(30)

	if test.Path == test.SecretKey {
		t.Fatalf("test Path is equal to test SecretKey")
	}
	if test.CreationDate == test.ExpirationDate {
		t.Fatalf("test creation and expiration date is same")
	}
	if test.Title != "Meditations" {
		t.Fatalf("title not set")
	}
	if test.Author != "Marcus Aurelius" {
		t.Fatalf("author not set")
	}
	if test.AuthorEmail != "marcus@rome.it" {
		t.Fatalf("email not set")
	}
	if test.TimeLimit != 3 {
		t.Fatalf("time limit is not right")
	}
	if test.AnswerSize != 4 {
		t.Fatalf("answer size is not right")
	}
}
