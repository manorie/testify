package storage

import (
	"github.com/manorie/testify/components/test"
	"testing"
)

func TestNewConnection(t *testing.T) {
	c := NewConnection()
	defer c.db.Close()

	err := c.db.Ping()
	if err != nil {
		t.Fatalf("Pinged database, returned error: %v", err)
	}
}

func TestWriteTest(t *testing.T) {
	test, _ := test.CreateNewTest()
	test.AddTitle("xyz")

	c := NewConnection()
	defer c.db.Close()

	c.WriteTest(test)
	if c.err != nil {
		t.Fatalf("Write caused an error: %v", c.err)
	}
}

func TestFindTestById(t *testing.T) {
	c := NewConnection()
	defer c.db.Close()

	test := c.FindTestById(1)
	if c.err != nil {
		t.Fatalf("Write caused an error: %v", c.err)
	}
	if test.Id != 1 {
		t.Fatalf("Find by Id did not return test with expected Id: %v", test.Id)
	}
}
