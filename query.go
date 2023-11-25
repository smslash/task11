package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getQueryString() {
	url := "http://www.example.com/search?query=blue%20shoes" // можно использовать `+` вместо `%20`

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Proto = "HTTP/1.1"

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
