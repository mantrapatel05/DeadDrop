package main

import (
	"context"
	"log"
	"net/http"

	"deadrop/config"
	"deadrop/server"
	"deadrop/store"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	ctx := context.Background()
	db, err := store.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	if err := db.Migrate(ctx); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Printf("server is running on http://localhost%s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, server.NewHandler(db)); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}
