package main

import (
	"net/http"

	"github.com/AlexanderT29/social-network-api/internal/store"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))

	app.store.Posts.Create(r.Context(), &store.Post{})
}
