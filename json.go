package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func postJSON() {
	userData := UserData{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	jsonData, err := json.Marshal(userData)
	if err != nil {
		log.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", "http://www.example.com/api/users", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Proto = "HTTP/1.1"

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(len(jsonData)))

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
