package server

import "net/http"

func (s *Server) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /", s.root)
	mux.HandleFunc("POST /drops", s.createDrop)
	mux.HandleFunc("GET /drops/{id}/status", s.status)
	mux.HandleFunc("POST /drops/{id}/knock", s.knock)
	mux.HandleFunc("GET /drops/{id}", s.getDrop)
	mux.HandleFunc("DELETE /drops/{id}", s.deleteDrop)
	mux.HandleFunc("PATCH /drops/{id}", s.patchDrop)
}
