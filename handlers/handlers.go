package handlers

import (
	"encoding/json"
	"go-chi-rest/posts"
	"log"
	"net/http"
)

func GetPosts(feed posts.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		if err := json.NewEncoder(w).Encode(items); err != nil {
			log.Fatal(err)
		}
	}
}

func CreatePost(feed posts.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rBody := map[string]string{}
		if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
			log.Fatal(err)
		}
		feed.Add(posts.Post{
			Title: rBody["title"],
			Text:  rBody["text"],
		})
		if _, err := w.Write([]byte("Success!")); err != nil {
			log.Fatal(err)
		}
	}
}
