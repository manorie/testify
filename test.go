package testify

import (
	"strconv"
	"time"
)

type Test struct {
	id             uint64
	title          string
	author         string
	path           string
	authorEmail    string
	isPublished    bool
	timeLimit      uint8
	answerSize     uint8
	creationDate   uint64
	expirationDate uint64
}

func CreateEmptyTest() (*Test, error) {
	return &Test{
		title: "alfa",
	}, nil
}

func ConstructASafePath() (string, error) {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	t := []rune(time)
	x := make([]rune, 10, 20)

	for _, r := range t {
		x = append(x, r)
		x = append(x, 'x')
	}
	return string(x), nil
}
