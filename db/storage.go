package storage

import (
	"database/sql"
	"github.com/manorie/testify/models/test"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func dbPath() string {
	if os.Getenv("testifyMode") == "production" {
		return "./productionDb.db"
	}
	return "./testDb.db"
}

func CreateTest(t *test.Test) error {
	db, err := sql.Open("sqlite3", dbPath())
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`INSERT INTO Tests(
        Title, Author, Path, SecretKey, AuthorEmail, IsPublished,
        TimeLimit, AnswerSize, CreationDate, ExpirationDate)
        values(?,?,?,?,?,?,?,?,?,?)`)

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

	return db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(
		&t.Id, &t.Title, &t.Author, &t.Path, &t.SecretKey,
		&t.AuthorEmail, &t.IsPublished, &t.TimeLimit, &t.AnswerSize,
		&t.CreationDate, &t.ExpirationDate)
}

//type Connection struct {
//db  *sql.DB
//err error
//}

//func NewConnection() (c Connection) {
//c.db, c.err = sql.Open("sqlite3", dbPath())
//return c
//}

//func (c *Connection) WriteTest(t *test.Test) {
//var stmt *sql.Stmt
//var res sql.Result
//var id int64

//stmt, c.err = c.db.Prepare(`INSERT INTO Tests(
//Title, Author, Path, SecretKey, AuthorEmail,
//IsPublished, TimeLimit, AnswerSize, CreationDate,
//ExpirationDate) values(?,?,?,?,?,?,?,?,?,?)`)

//res, c.err = stmt.Exec(&t.Title, &t.Author, &t.Path, &t.SecretKey,
//&t.AuthorEmail, &t.IsPublished, &t.TimeLimit, &t.AnswerSize,
//&t.CreationDate, &t.ExpirationDate)

//id, c.err = res.LastInsertId()

//c.err = c.db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
//&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
//&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)

//}

//func (c *Connection) FindTestById(id int) (t test.Test) {
//c.err = c.db.QueryRow("SELECT * FROM Tests WHERE Id = ?", id).Scan(&t.Id,
//&t.Title, &t.Author, &t.Path, &t.SecretKey, &t.AuthorEmail, &t.IsPublished,
//&t.TimeLimit, &t.AnswerSize, &t.CreationDate, &t.ExpirationDate)

//return t
/*}*/
