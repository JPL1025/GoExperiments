package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Prints Hello World
func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello World! ")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World!")
	case "/x":
		fmt.Fprint(w, "x!")
	default:
		fmt.Fprint(w, "Error!")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func helloWorldPageDark(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello (Dark) World!\n")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello (Dark) World!</h1>")
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	//w.Header().Set("Content-Type", "text/plain") // demonstrate plaintext
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Attempted timeout")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Didn't timeout")
}

// URL Handler
func url(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	if name == "" {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "(Url) Hello %s!", name)
}

// Body Handler
func body(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	name, ok := data["name"]
	if !ok {
		http.Error(w, "name is empty", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "(Body) Hello %s!", name)
}

// {"name": "xml1025"}
func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
		// http://localhost/url?name=xml1025
	case "/body":
		body(w, r)
		// http://localhost/body
		// with body in json: {"name": "xml1025"}
	default:
		w.Write([]byte("(Default) Hello World!"))
	}
}

func main() {
	http.HandleFunc("/", handle)
	//http.HandleFunc("/time", timeout)
	http.ListenAndServe("", nil)

	/*
		server := http.Server{
			Addr:         "",
			Handler:      nil,
			ReadTimeout:  1000,
			WriteTimeout: 1000,
		}

		var muxDarkMode http.ServeMux
		server.Handler = &muxDarkMode
		muxDarkMode.HandleFunc("/", helloWorldPageDark)

		server.ListenAndServe()
	*/
}
