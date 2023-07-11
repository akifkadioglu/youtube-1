package route

import (
	"github.com/akifkadioglu/youtube-1/pkg/auth"
	"github.com/akifkadioglu/youtube-1/pkg/home"
	"github.com/akifkadioglu/youtube-1/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func CreateServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/index", home.Index)
	r.Post("/register", auth.Register)
	r.Post("/login", auth.Login)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth()))
		r.Use(jwtauth.Authenticator)

		r.Get("/", home.Index)
	})
	return r
}
