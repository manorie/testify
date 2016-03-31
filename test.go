package testify

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Test struct {
	Id             uint64
	Title          string
	Author         string
	Path           string
	SecretKey      string
	AuthorEmail    string
	IsPublished    bool
	TimeLimit      uint8
	AnswerSize     uint8
	CreationDate   int64
	ExpirationDate int64
}

func (t *Test) AddAuthor(author string) error {
	return t.setString(&t.Author, author)
}

func (t *Test) setString(p *string, v string) error {
	if len(v) > 255 {
		return fmt.Errorf("field %s is too long.", v)
	}
	*p = v
	return nil
}

func CreateNewTest() (*Test, error) {
	return &Test{
		Path:         constructAUniqueKey(),
		SecretKey:    constructAUniqueKey(),
		IsPublished:  false,
		CreationDate: time.Now().Unix(),
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz-_~")
var tempRune = make([]rune, 20, 20)

func constructAUniqueKey() string {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	t := []rune(now)
	x := make([]rune, 20, 20)

	ns := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(ns)

	for _, r := range t {
		x = append(x, letterRunes[r1.Intn(len(letterRunes))])
		x = append(x, r)
	}
	return string(x)
}
