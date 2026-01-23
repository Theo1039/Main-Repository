package main

import (
	"log"
	"net/http"
	"t-products/cars"
	"t-products/electronics"
	"t-products/phones"
)

func main() {
	http.HandleFunc("/cars", cars.Handler)
	http.HandleFunc("/electronics", electronics.Handler)
	http.HandleFunc("/phones", phones.Handler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
