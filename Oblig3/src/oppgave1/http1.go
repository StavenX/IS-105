package main

import (
	"io"
	"net/http"
)

// Simple method that greets when accessing the server
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello client.")
}

func main() {
	// If path that is accessed is "/", server will respond with 'hello' method
	http.HandleFunc("/", hello)
	// Server listens on port 8080
	http.ListenAndServe(":8080", nil)
}