package main

import (
	"github.com/elkcityhazard/contact-keeper/cmd/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(StripTrailingSlash)
	mux.Use(DefaultHeaders)

	mux.Get("/ping", handlers.Repo.PingHandler)

	mux.Get("/", handlers.Repo.HomeHandler)

	return mux
}
