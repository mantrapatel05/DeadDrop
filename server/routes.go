package server

import "net/http"

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", health)
}