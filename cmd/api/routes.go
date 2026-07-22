package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORSE)

	mux.Get("/", app.Hello)
	mux.Get("/movies", app.movies)
	mux.Get("/authenticate", app.authenticate)
	return mux
}
