package main

import (
	"fmt"
	"github.com/manorie/testify/models/test"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, e := test.BuildEmptyTest()

	t.AddTitle(r.FormValue("title"), e)
	t.AddAuthor(r.FormValue("author"), e)
	t.AddAuthorEmail(r.FormValue("author_email"), e)
	t.SetTimeLimit(parseUI8(r.FormValue("time_limit")), e)
	t.SetAnswerSize(parseUI8(r.FormValue("answer_size")), e)
	t.SetExpireInDays(parseUI8(r.FormValue("expire_in")), e)

	if e.NotEmpty() {
		fmt.Fprintf(w, e.ToJson())
	} else {
		fmt.Fprintf(w, t.ToJson())
	}
}

func parseUI8(s string) uint8 {
	i, e := strconv.ParseUint(s, 10, 8)
	if e != nil {
		return uint8(0)
	}
	return uint8(i)
}

func main() {
	http.HandleFunc("/create_test", handler)
	http.ListenAndServe(":8080", nil)
}
