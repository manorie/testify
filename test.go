package testify

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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

func (t *Test) WriteTo(db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO Tests(
        Title, Author, Path, SecretKey, AuthorEmail,
        IsPublished, TimeLimit, AnswerSize, CreationDate,
        ExpirationDate) values(?,?,?,?,?,?,?,?,?,?)`)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(&t.Title, &t.Author, &t.Path, &t.SecretKey,
		&t.AuthorEmail, &t.IsPublished, &t.TimeLimit, &t.AnswerSize,
		&t.CreationDate, &t.ExpirationDate)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	return db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
		&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
		&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)
}

func FindTestById(db *sql.DB, id int) (Test, error) {
	var t Test
	err := db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
		&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
		&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)

	return t, err
}

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/testifyDb.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Tests (
            Id              integer PRIMARY KEY AUTOINCREMENT,
            Title           varchar(255),
            Author          varchar(255),
            Path            varchar(255) NOT NULL,
            SecretKey       varchar(255) NOT NULL,
            AuthorEmail     varchar(255) NOT NULL,
            IsPublished     boolean,
            TimeLimit       integer NOT NULL,
            AnswerSize      integer NOT NULL,
            CreationDate    integer NOT NULL,
            ExpirationDate  integer NOT NULL)`)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS TestPathsIndex
            ON Tests(path)`)

	if err != nil {
		panic(err)
	}
	return db, err
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
