package main

import (
	"log"
	"net/http"
	"short-link/middleware"
	"short-link/repository"
	"short-link/router"
)

func main() {
	repository.ConnectDB();
	
	mux := router.ShortRoutes();

	handleWithCors := middleware.CorsMiddleware(mux)
	
	log.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", handleWithCors)
}