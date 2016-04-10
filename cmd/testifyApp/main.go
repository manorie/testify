package main

import (
	"encoding/json"
	"fmt"
	"github.com/manorie/testify/models/test"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	test := test.Test{
		Title: r.FormValue("title"),
	}

	j, _ := json.Marshal(test)

	fmt.Fprintf(w, string(j))
}

func main() {
	http.HandleFunc("/create_test", handler)
	http.ListenAndServe(":8080", nil)
}
