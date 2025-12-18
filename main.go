package main

import (
	"log"
	"net/http"

	api "vercel-envs/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
