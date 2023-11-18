package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// struct{}{} define an struct and instanciate it as empty struct.
	respondWithJson(w, 200, struct{}{})
}
