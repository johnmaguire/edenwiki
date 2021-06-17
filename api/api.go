package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	h := Handlers{newDatabase()}
	r.Put("/page", h.putPage)

	return r
}
