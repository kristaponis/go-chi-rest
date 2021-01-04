package main

import (
	"github.com/go-chi/chi"
	"go-chi-rest/handlers"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handlers.PostsHandler)
	r.Post("/", handlers.PostsHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
