package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Simpler POST Request")

	const myPostApiUrl = "https://jsonplaceholder.typicode.com/posts"

	// Fake JSON payload
	requestBody := strings.NewReader(`{
		"title": "foo",
		"body": "bar",
		"userId": 1
	}`)

	// Make POST request
	response, err := http.Post(myPostApiUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Defer only after successful request

	// Read response
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response:")
	fmt.Println(string(dataBytes))
}
