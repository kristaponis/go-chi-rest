package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-chi-rest/handlers"
	"go-chi-rest/posts"
	"log"
	"net/http"
)

const (
	PORT = ":8080"
)

var (
	feed = posts.New()
)

func main() {
	r := chi.NewRouter()

	r.Get("/posts", handlers.GetPosts(feed))
	r.Post("/posts", handlers.CreatePost(feed))

	fmt.Println("Serving application...")
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}

