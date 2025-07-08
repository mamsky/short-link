package main

import (
	"log"
	"net/http"
	"short-link/repository"
	"short-link/router"
)

func main() {
	repository.ConnectDB()

	router.ShortRoutes()

	log.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}