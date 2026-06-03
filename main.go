package main

import (
	"net/http"
	"log"

	"deadrop/server"
)

func main(){
	port := ":8080"
	log.Printf("server is running on http://localhost%s",port)
	if err := http.ListenAndServe(port, server.NewHandler()); err != nil {
		log.Println("server is stopped with error:",err)
		return
	}
}