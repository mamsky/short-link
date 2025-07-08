package router

import (
	"net/http"
	"short-link/handler"
)

func ShortRoutes() {
	http.HandleFunc("/", handler.HandleLongUrl)
}