package main

import (
	"fmt"
	"net/http"
)

// {"name": "xml1025"}
func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": // TODO
	case "/body": // TODO
	default:

		fmt.Fprintf(w, "Hello, World!")
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("", nil)
}
