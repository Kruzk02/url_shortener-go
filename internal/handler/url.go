package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"url_shortener/internal/dto"
	"url_shortener/internal/service"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 10)

func GetOriginByCodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code is required", http.StatusBadRequest)
		return
	}
	println(limiter.Tokens())
	if !limiter.Allow() {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	url, err := service.GetOriginByCode(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	http.Redirect(w, r, url.Origin, http.StatusSeeOther) // Test on web browser instead tools like postman
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SaveUrlHandler(w http.ResponseWriter, r *http.Request) {
	var urlDTO dto.UrlDTO

	if err := json.NewDecoder(r.Body).Decode(&urlDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	println(limiter.Tokens())

	if !limiter.Allow() {
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	code, err := service.Save(urlDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reponse := map[string]string{"code": code}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(reponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
