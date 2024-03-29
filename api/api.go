package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/johnmaguire/edenwiki/git"
	"github.com/rs/cors"
)

type handlers struct {
	wiki *git.Wiki
}

func NewRouter(wiki *git.Wiki) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:8080"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}).Handler)

	h := handlers{wiki}
	r.Get("/page", h.listPages)
	r.Put("/page/{pageName}", h.putPage)
	r.Get("/page/{pageName}", h.getPage)

	return r
}
