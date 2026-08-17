package main

import (
	"fmt"
	"net/http"
)
func main() {

	// Serve all files in current folder
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}