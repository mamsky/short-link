package main

import (
	"log"
	"net/http"
	"short-link/repository"
)

func main() {
	repository.ConnectDB()

	log.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}