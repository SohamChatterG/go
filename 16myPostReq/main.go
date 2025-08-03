// package main: Every executable Go program must start with this package.
package main

// Import necessary packages from the Go standard library.
import (
	"bytes"         // For creating an io.Reader from a byte slice.
	"encoding/json" // For marshaling Go structs into JSON.
	"fmt"           // For formatted I/O (printing to the console).
	"io"            // For I/O primitives, like reading a stream.
	"log"           // For simple logging and fatal error handling.
	"net/http"      // For making HTTP requests.
	"time"          // For setting timeouts on our HTTP client.
)

// Dummy API endpoint for POST requests.
const postAPIURL = "https://jsonplaceholder.typicode.com/posts"

// PostData represents the structure of the data we will send in the POST request body.
// The `json:"..."` tags are crucial for the `json.Marshal` function.
type PostData struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// ResponseData represents the structure of the data we expect to receive
// back from the API after a successful POST request.
type ResponseData struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// The main function is the entry point of the program.
func main() {
	fmt.Println("--- GoLang POST Request Tutorial ---")

	// 1. Prepare the data to be sent.
	// Create a Go struct with the data you want to send in the request body.
	requestBody := PostData{
		Title:  "GoLang Post Request",
		Body:   "This is a test post from a Go application.",
		UserID: 1,
	}

	// 2. Marshal the Go struct into a JSON byte slice.
	// We must convert our struct into a byte slice because that's what the
	// HTTP request body expects.
	postBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}

	// 3. Create an HTTP client with best practices.
	// It's a best practice to create and reuse a client with a timeout.
	// The timeout prevents the request from hanging indefinitely.
	client := &http.Client{Timeout: 10 * time.Second}

	// 4. Create the POST request.
	// http.NewRequest gives you full control over the request.
	// We pass the URL, the method ("POST"), and the request body as an `io.Reader`.
	// `bytes.NewBuffer` creates an `io.Reader` from our JSON byte slice.
	req, err := http.NewRequest("POST", postAPIURL, bytes.NewBuffer(postBodyBytes))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 5. Set the necessary headers.
	// This is crucial for most APIs. The `Content-Type` header tells the server
	// what kind of data to expect in the request body.
	req.Header.Set("Content-Type", "application/json")

	// 6. Execute the request.
	// The `client.Do(req)` method sends the request and waits for a response.
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}

	// 7. Ensure the response body is closed.
	// This is a critical Go idiom. The `defer` keyword ensures this call
	// happens just before the function returns, even if an error occurs.
	defer response.Body.Close()

	// 8. Check the HTTP status code.
	// A successful POST request typically returns a 201 Created status.
	if response.StatusCode != http.StatusCreated {
		log.Fatalf("API returned a non-success status code: %s", response.Status)
	}

	// 9. Read the entire response body.
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 10. Unmarshal the JSON response into our Go struct.
	// This is the final step to convert the JSON data back into a usable Go type.
	var responseData ResponseData
	err = json.Unmarshal(responseBodyBytes, &responseData)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	// 11. Print the result.
	// We can now access the fields of our `responseData` struct.
	fmt.Println("\n--- API Response ---")
	fmt.Printf("Received response with ID: %d\n", responseData.ID)
	fmt.Printf("Title: %s\n", responseData.Title)
	fmt.Printf("Body: %s\n", responseData.Body)
}
