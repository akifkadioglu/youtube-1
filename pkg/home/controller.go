package home

import (
	"net/http"

	"github.com/go-chi/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAlreadyReported)
	render.JSON(w, r, map[string]string{
		"hello": "world",
	})
}
