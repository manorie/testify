package storage

import (
	"database/sql"
	"github.com/manorie/testify/components/test"
	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	db  *sql.DB
	err error
}

func NewConnection() (c Connection) {
	c.db, c.err = sql.Open("sqlite3", "./testifyDb.db")

	_, c.err = c.db.Exec(`CREATE TABLE IF NOT EXISTS Tests (
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

	_, c.err = c.db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS TestPathsIndex
            ON Tests(path)`)

	return c
}

func (c *Connection) WriteTest(t *test.Test) {
	var stmt *sql.Stmt
	var res sql.Result
	var id int64

	stmt, c.err = c.db.Prepare(`INSERT INTO Tests(
        Title, Author, Path, SecretKey, AuthorEmail,
        IsPublished, TimeLimit, AnswerSize, CreationDate,
        ExpirationDate) values(?,?,?,?,?,?,?,?,?,?)`)

	res, c.err = stmt.Exec(&t.Title, &t.Author, &t.Path, &t.SecretKey,
		&t.AuthorEmail, &t.IsPublished, &t.TimeLimit, &t.AnswerSize,
		&t.CreationDate, &t.ExpirationDate)

	id, c.err = res.LastInsertId()

	c.err = c.db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
		&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
		&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)

}

func (c *Connection) FindTestById(id int) (t test.Test) {
	c.err = c.db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
		&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
		&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)

	return t
}
