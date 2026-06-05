package server

import (
	"deadrop/store"
	"net/http"
)

type Server struct {
	db *store.DB
}

func NewHandler(db *store.DB) http.Handler {
	s := &Server{db: db}
	mux := http.NewServeMux()
	s.registerRoutes(mux)
	return mux
}
