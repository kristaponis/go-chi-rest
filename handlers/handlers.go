package handlers

import (
	"encoding/json"
	"go-chi-rest/posts"
	"net/http"
)

var Feed = posts.New()

func GetPosts(w http.ResponseWriter, r *http.Request) {
	items := Feed.GetAll()
	json.NewEncoder(w).Encode(items)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	rBody := map[string]string{}
	json.NewDecoder(r.Body).Decode(&rBody)
	Feed.Add(posts.Post{
		Title: rBody["title"],
		Text:  rBody["text"],
	})

	w.Write([]byte("Good!"))
}
