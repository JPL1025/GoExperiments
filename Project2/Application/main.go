package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "Project2/Application/login.html"

	temp, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error Parsing File", err)
		return
	}
	err = temp.Execute(w, nil)
	if err != nil {
		fmt.Println("Error Executing Template", err)
		return
	}
}

// Obviously not real since plaintext is a terrible way to store passwords
var userDatabase = map[string]string{
	"xml1025": "demoPassword",
}

// Obviously not a real login since it would be terrible to store a password in plaintext
func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println(username, password)

	if userDatabase[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Printf("User Successfully Logged In ")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("Valid Credential Not Found ")
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit": // TODO
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Hello, World!")
	}
}

func main() {
	http.HandleFunc("/", handle)
	//http.ListenAndServe("", nil)
	http.ListenAndServeTLS("", "Project2/Application/cert.pem", "Project2/Application/key.pem", nil)
}
