package main

import (
	"log"
	"net/http"

	"backend/api"
)

func main() {
	http.HandleFunc("/search", api.SearchHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
