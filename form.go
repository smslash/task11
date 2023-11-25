package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func postForm() {
	data := strings.NewReader("username=admin&password=secret")

	req, err := http.NewRequest("POST", "http://www.example.com/login", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Proto = "HTTP/1.1"

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(body))
}
