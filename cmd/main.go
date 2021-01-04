package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-chi-rest/handlers"
	"log"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {
	r := chi.NewRouter()

	r.Get("/posts", handlers.PostsHandler)
	r.Post("/posts", handlers.PostsHandler)

	fmt.Println("Serving application...")
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
