package main

import (
	"net/http"
	"log"

	"deadrop/config"
	"deadrop/server"
)

func main(){
	cfg,err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Printf("server is running on http://localhost%s",cfg.Port)
	if err := http.ListenAndServe(cfg.Port, server.NewHandler()); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}