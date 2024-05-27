package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	// Create the JSON object
	reqData := map[string]string{
		"name": "xml1025",
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost/body", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	response, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		response.Body.Close()
	}
}
