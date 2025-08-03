package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(" Web Req")

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type in Go %T\n", response)
	fmt.Println(response)
	// Understand TCP pooling to understand why go does not close the request body
	// TCP pooling is the practice of reusing existing TCP connections instead of creating a new connection for every request.
	// This is especially useful when multiple HTTP requests are made to the same server/host.

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}

/*
 Why Go Doesn't Automatically Close response.Body
🧠 1. HTTP Keep-Alive (Persistent Connections)
Go follows the HTTP/1.1 spec where connections are kept alive by default.

This avoids the overhead of repeated TCP handshakes for every request to the same host.

Keeping the TCP connection open allows Go to reuse it for future requests, improving performance.

📦 2. Response is a Stream
Go returns the response body as an io.ReadCloser stream, not as a fully loaded string or object.

Go doesn’t assume when or how much you want to read — you control that.

As long as the stream (i.e., response.Body) is open, the underlying connection remains in use.

✋ 3. You Might Not Read the Whole Body
If you don’t read the full body and don’t call Close(), Go can’t reuse the connection.

If you only read part of the stream, you must still Close() it to tell Go, "I'm done."

🗑️ 4. Resource Cleanup is Your Responsibility
Not closing the body leaks:

Memory

File descriptors

Network sockets

Over time, this leads to failures like too many open files.

💡 5. Explicit Control Matches Go's Design
Go avoids hidden magic.

You explicitly read the response and close it — very C-like in spirit.

This gives you low-level control over how data is processed.
*/
