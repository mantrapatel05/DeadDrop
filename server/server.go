package server

import (
	"deadrop/store"
	"net/http"
)

type Server struct {
	db *store.DB
	encryptionKey []byte
}

func NewHandler(db *store.DB, encryptionKey []byte) http.Handler {
	s := &Server{db: db, encryptionKey: encryptionKey}
	mux := http.NewServeMux()
	s.registerRoutes(mux)
	return mux
}
