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

func (t *Test) AddAuthor(name string) (*Test, error) {
	if err := validateStringLength(name); err != nil {
		return nil, err
	}
	t.Author = name
	return t, nil
}

func (t *Test) AddTitle(title string) (*Test, error) {
	if err := validateStringLength(title); err != nil {
		return nil, err
	}
	t.Title = title
	return t, nil
}

func validateStringLength(field string) error {
	if len(field) > 255 {
		return fmt.Errorf("field %v is too long.", field)
	}
	return nil
}

func CreateNewTest() (*Test, error) {
	path, _ := constructAUniqueKey()
	secretKey, _ := constructAUniqueKey()

	return &Test{
		Path:         path,
		SecretKey:    secretKey,
		IsPublished:  false,
		CreationDate: time.Now().Unix(),
	}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz-_~")

func constructAUniqueKey() (string, error) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	t := []rune(now)
	x := make([]rune, 20, 20)

	ns := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(ns)

	for _, r := range t {
		x = append(x, letterRunes[r1.Intn(len(letterRunes))])
		x = append(x, r)
	}
	return string(x), nil
}
