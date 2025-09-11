package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/welcome", welcomePage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Store !"))
	headers := r.Header
	fmt.Println("Request Headers:")
	for key, values := range headers {
		fmt.Printf("%s: %v\n", key, values)
	}
	method := r.Method
	path := r.URL.Path
	fmt.Printf("HTTP Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
}
