package router

import (
	"net/http"
	"short-link/handler"
)

func ShortRoutes() http.Handler{
	mux := http.NewServeMux();
	mux.HandleFunc("/shorten", handler.HandleLongUrl)
	mux.HandleFunc("/", handler.HandleRedirect)
	return mux
}