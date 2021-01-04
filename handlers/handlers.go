package handlers

import (
	"encoding/json"
	"go-chi-rest/posts"
	"net/http"
)

var feed = posts.New()

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	items := feed.GetAll()
	json.NewEncoder(w).Encode(items)
}
