package route

import (
	"github.com/akifkadioglu/youtube-1/pkg/home"
	"github.com/akifkadioglu/youtube-1/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func CreateServer() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/asd", home.Index)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth()))
		r.Use(jwtauth.Authenticator)
		
		r.Get("/",home.Index)
	})
	return r
}
