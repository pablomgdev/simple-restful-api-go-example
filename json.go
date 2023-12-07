package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// interface{} means any type.
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

// TODO: continue here: https://youtu.be/un6ZyFkqFKo?si=Q66UPPoKTJQTgcYk&t=25598
func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{
		Error: message,
	})
}
