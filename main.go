package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/siddhant-vij/Auth0-API-Demo/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	rtr := router.New()

	log.Fatal(http.ListenAndServe(":8080", rtr))
}
