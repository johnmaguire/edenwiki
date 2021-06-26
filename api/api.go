package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/johnmaguire/gardenwiki/api/data"
)

type Handlers struct {
	db *database
}

type database struct {
	Pages map[string]data.Page
}

func newDatabase() *database {
	return &database{map[string]data.Page{}}
}

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))

	h := Handlers{newDatabase()}
	r.Get("/page", h.listPages)
	r.Put("/page/{pageName}", h.putPage)

	return r
}
