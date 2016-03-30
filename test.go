package testify

import (
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
