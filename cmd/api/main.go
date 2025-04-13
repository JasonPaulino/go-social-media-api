package main

import (
	"log"
	"net/http"

	"github.com/JasonPaulino/go-social-media-api/api/v1"
)

func main() {
	router := v1.Routes()

	log.Println("Server is listening on port :8080")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
	}
}
