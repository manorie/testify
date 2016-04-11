package main

import (
	"encoding/json"
	"fmt"
	"github.com/manorie/testify/models/test"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	test := test.BuildEmptyTest()

	err := test.AddTitle(r.FormValue("title"))
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	err = test.AddAuthorEmail(r.FormValue("author_email"))
	if err != nil {
		fmt.Fprintf(w, errJson(err.Error()))
		return
	}

	j, _ := json.Marshal(test)
	fmt.Fprintf(w, string(j))
}

func errJson(err string) string {
	return `{error: "` + err + `"}`
}

func main() {
	http.HandleFunc("/create_test", handler)
	http.ListenAndServe(":8080", nil)
}
