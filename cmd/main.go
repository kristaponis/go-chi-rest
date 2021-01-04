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

func main() {
	handlers.Feed.Add(posts.Post{
		Title: "First post",
		Text: "This is some nice text",
	})

	r := chi.NewRouter()

	r.Get("/posts", handlers.GetPosts)
	r.Post("/posts", handlers.CreatePost)

	fmt.Println("Serving application...")
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
