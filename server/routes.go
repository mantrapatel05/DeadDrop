package server

import "net/http"

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /", s.root)
	mux.HandleFunc("POST /drops", s.createDrop)
}