package main

import (
	"log"
	"net/http"
	"url_shortener/api"
	"url_shortener/internal/repository"
)

func main() {
	repository.InitDB()
	router := api.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8000", router))
}
