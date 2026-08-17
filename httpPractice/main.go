package main

import (
	"fmt"
	"net/http"
)

func add(a int, b int) int {
	return a + b
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To Theo App")
}

func addition(w http.ResponseWriter, r *http.Request) {
	ad := add(3, 5)
	fmt.Fprintln(w, ad)
}
func htmlPractice(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/add", addition)
	http.HandleFunc("/html", htmlPractice)

	fmt.Println("server run at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}