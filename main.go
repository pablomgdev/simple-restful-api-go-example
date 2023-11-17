package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("There was an error loading the environment variables.")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("There was an error getting the port number.")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options {
		AllowedOrigins:		[]string{"https://*", "http://*"},
		AllowedMethods:		[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:		[]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:		[]string{"Link"},
		AllowCredentials:	false,
		MaxAge:						300,
	}))
	server := &http.Server {
		Handler: router,
		Addr: ":" + portString,
	}
	fmt.Printf("Starting server on port %v", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
