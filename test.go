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

func (t *Test) AddTitle(title string) error {
	return t.setString(&t.Title, title, false)
}

func (t *Test) AddAuthor(author string) error {
	return t.setString(&t.Author, author, false)
}

func (t *Test) AddAuthorEmail(email string) error {
	return t.setString(&t.AuthorEmail, email, true)
}

func (t *Test) SetTimeLimit(time uint8) error {
	return t.setUint8(&t.TimeLimit, time, 1, 30)
}

func (t *Test) SetAnswerSize(size uint8) error {
	return t.setUint8(&t.AnswerSize, size, 2, 5)
}

func (t *Test) Publish() error {
	if t.IsPublished {
		return fmt.Errorf("Test is already published.")
	}
	t.IsPublished = true
	return nil
}

func (t *Test) UnPublish() error {
	if !t.IsPublished {
		return fmt.Errorf("Test is already unpublished.")
	}
	t.IsPublished = false
	return nil
}

func (t *Test) SetExpireInDays(days uint8) error {
	if t.CreationDate == 0 {
		fmt.Errorf("can't set expiration date, creation date of test is null")
	}
	if days > 180 {
		return fmt.Errorf("expire in days %d is above max limit of %d", days, 180)
	}
	if days < 1 {
		return fmt.Errorf("expire in days %d is below min limit of %d", days, 1)
	}
	c := time.Unix(t.CreationDate, 0).AddDate(0, 0, int(days))
	t.ExpirationDate = c.Unix()

	return nil
}

func (t *Test) setUint8(p *uint8, v uint8, min uint8, max uint8) error {
	if v > max {
		return fmt.Errorf("time field %d is above max limit of %d", v, max)
	}
	if v < min {
		return fmt.Errorf("time field %d is below min limit of %d", v, min)
	}
	*p = v
	return nil
}

func (t *Test) setString(p *string, v string, notNull bool) error {
	if len(v) > 255 {
		return fmt.Errorf("field %s is too long.", v)
	}
	if len(v) == 0 && notNull {
		return fmt.Errorf("field %s is empty.", v)
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
