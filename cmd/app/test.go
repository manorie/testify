package main

import (
	"fmt"
	"time"
)

type Test struct {
	Title          string     `json:"title"`
	Author         string     `json:"author"`
	Path           string     `json:"path"`
	SecretKey      string     `json:"-"`
	AuthorEmail    string     `json:"author_email"`
	IsPublished    bool       `json:"is_published"`
	TimeLimit      uint8      `json:"time_limit"`
	AnswerSize     uint8      `json:"answer_size"`
	CreationDate   *time.Time `json:"creation_date,omitempty"`
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
}

type TestPost struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	AuthorEmail string `json:"author_email"`
	IsPublished bool   `json:"is_published"`
	TimeLimit   uint8  `json:"time_limit"`
	AnswerSize  uint8  `json:"answer_size"`
	ExpireIn    uint8  `json:"expire_in"`
}

type TestErrors struct {
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	AuthorEmail string `json:"author_email,omitempty"`
	IsPublished string `json:"is_published,omitempty"`
	TimeLimit   string `json:"time_limit,omitempty"`
	AnswerSize  string `json:"answer_size,omitempty"`
	ExpireIn    string `json:"expire_in,omitempty"`
	ErrorCount  uint8  `json:"error_count"`
}

func init() {
	p := TestPost{Title: "Alfa"}
	err, t := p.New()
	fmt.Println(err)
	fmt.Println(t)
}

func (post *TestPost) New() (TestErrors, Test) {
	e := TestErrors{}

	e.ValStringPresence(&e.Title, &post.Title, false)
	e.ValStringPresence(&e.Author, &post.Author, false)
	e.ValStringPresence(&e.AuthorEmail, &post.AuthorEmail, true)

	return e, Test{}
}

func (e *TestErrors) IncrementErrorCount() {
	e.ErrorCount += 1
}

func (te *TestErrors) ValStringPresence(eF *string, p *string, notNull bool) {
	if len(*p) > 255 {
		*eF = "field is too long."
		te.IncrementErrorCount()
	}
	if len(*p) == 0 && notNull {
		*eF = "field can't be empty."
		te.IncrementErrorCount()
	}
}
