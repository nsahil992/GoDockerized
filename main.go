// main.go
package main

import (
    "fmt"
    "net/http"
)

// handler function will respond with "Hello, World!"
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    // Register the handler for the root URL "/"
    http.HandleFunc("/", handler)

    // Print a message indicating the server is starting
    fmt.Println("Server starting on port 8080...")

    // Start the server on port 8080
    http.ListenAndServe(":8080", nil)
}

