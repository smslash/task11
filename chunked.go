package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func transferEncoding() {
	reader, writer := io.Pipe()

	go func() {
		time.Sleep(1 * time.Second)
		chunks := []string{"Hello, ", "world!"}
		for _, chunk := range chunks {
			time.Sleep(1 * time.Second)
			if _, err := writer.Write([]byte(chunk)); err != nil {
				log.Fatal(err)
			}
		}
		writer.Close()
	}()

	req, err := http.NewRequest("POST", "http://www.example.com/api/chunked", reader)
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
