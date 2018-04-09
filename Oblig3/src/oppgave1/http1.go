package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello client.")
}

func test(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello test .")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/1", test)
	http.ListenAndServe(":8080", nil)
}