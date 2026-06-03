package server

import "net/http"

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}