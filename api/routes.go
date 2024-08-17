package api

import (
	"url_shortener/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/url", handler.GetOriginByCodeHandler).Methods("GET")
	r.HandleFunc("/url", handler.SaveUrlHandler).Methods("POST")
	return r
}
