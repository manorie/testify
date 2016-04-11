package test

import (
	"encoding/json"
	"fmt"
	"github.com/manorie/testify/lib/keyMaker"
	"time"
)

type Test struct {
	Id             uint64 `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Path           string `json:"path"`
	SecretKey      string `json:"secret_key"`
	AuthorEmail    string `json:"author_email"`
	IsPublished    bool   `json:"is_published"`
	TimeLimit      uint8  `json:"time_limit"`
	AnswerSize     uint8  `json:"answer_size"`
	CreationDate   int64  `json:"creation_date"`
	ExpirationDate int64  `json:"expiration_date"`
}

type TestErrors struct {
	Errors map[string][]string `json:"errors"`
}

func (t *TestErrors) NotEmpty() bool {
	return len(t.Errors) != 0
}

func (t *TestErrors) ToJson() string {
	j, _ := json.Marshal(t)
	return string(j)
}

func (t *Test) ToJson() string {
	j, _ := json.Marshal(t)
	return string(j)
}

func BuildEmptyTest() (*Test, *TestErrors) {
	return &Test{
		Path:         keyMaker.Generate(),
		SecretKey:    keyMaker.Generate(),
		IsPublished:  false,
		CreationDate: time.Now().Unix(),
	}, &TestErrors{make(map[string][]string)}
}

func (e *TestErrors) CheckError(title string, err error) {
	if err != nil {
		e.Errors[title] = append(e.Errors[title], err.Error())
	}
}

func (t *Test) AddTitle(title string, e *TestErrors) {
	err := t.setString(&t.Title, title, false)
	e.CheckError("title", err)
}

func (t *Test) AddAuthor(author string, e *TestErrors) {
	err := t.setString(&t.Author, author, false)
	e.CheckError("author", err)
}

func (t *Test) AddAuthorEmail(email string, e *TestErrors) {
	err := t.setString(&t.AuthorEmail, email, true)
	e.CheckError("author_email", err)
}

func (t *Test) SetTimeLimit(time uint8, e *TestErrors) {
	err := t.setUint8(&t.TimeLimit, time, 1, 30)
	e.CheckError("time_limit", err)
}

func (t *Test) SetAnswerSize(size uint8, e *TestErrors) {
	err := t.setUint8(&t.AnswerSize, size, 2, 5)
	e.CheckError("answer_size", err)
}

func (t *Test) setString(p *string, v string, notNull bool) error {
	if len(v) > 255 {
		return fmt.Errorf("field is too long.")
	}
	if len(v) == 0 && notNull {
		return fmt.Errorf("field is empty.")
	}
	*p = v
	return nil
}

func (t *Test) setUint8(p *uint8, v uint8, min uint8, max uint8) error {
	if v > max {
		return fmt.Errorf("field is above max limit of %d.", max)
	}
	if v < min {
		return fmt.Errorf("field is below min limit of %d.", min)
	}
	*p = v
	return nil
}

func (t *Test) Publish(e *TestErrors) {
	if t.IsPublished {
		e.CheckError("is_published", fmt.Errorf("Test is already published."))
	}
	t.IsPublished = true
}

func (t *Test) UnPublish(e *TestErrors) {
	if !t.IsPublished {
		e.CheckError("is_published", fmt.Errorf("Test is already unpublished."))
	}
	t.IsPublished = false
}

func (t *Test) SetExpireInDays(days uint8, e *TestErrors) {
	err := t.setExpireInDays(days)
	e.CheckError("expire_in", err)
}

func (t *Test) setExpireInDays(days uint8) error {
	if t.CreationDate == 0 {
		return fmt.Errorf("can't set expiration date, creation date of test is null")
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
