package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	// struct{}{} define an struct and instanciate it as empty struct.
	respondWithError(w, 400, "Something went wrong")
}
