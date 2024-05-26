package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Plain URL method - http://localhost
	response, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// URL key-value form - http://localhost/url?name=xml1025
	response, err = http.Get("http://localhost/url?name=xml1025")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// JSON body method - http://localhost/body
	client := &http.Client{}
	reqBody := `{"name":"xml1025"}`
	req, err := http.NewRequest("POST", "http://localhost/body", strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}
}
