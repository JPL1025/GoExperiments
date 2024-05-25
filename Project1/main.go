package main

import (
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

func main() {
	//http.HandleFunc("/", htmlVsPlain)
	http.HandleFunc("/time", timeout)
	//http.ListenAndServe("", nil)

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
}
