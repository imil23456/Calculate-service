package main

import (
	"calculator/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", handler.HandleCalculation)

	log.Println("Server is running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
