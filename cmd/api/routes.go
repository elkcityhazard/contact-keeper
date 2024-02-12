package main

import (
	"net/http"

	"github.com/elkcityhazard/contact-keeper/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(StripTrailingSlash)
	mux.Use(DefaultHeaders)
	mux.Use(app.SessionManager.LoadAndSave)

	mux.Mount("/api/v1", apiRoutes())
	mux.Mount("/api/v1/users", apiUserRoutes())

	return mux
}

func apiRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(StripTrailingSlash)
	router.Use(DefaultHeaders)

	router.Get("/ping", handlers.Repo.PingHandler)
	router.Get("/", handlers.Repo.HomeHandler)

	return router
}

func apiUserRoutes() http.Handler {
	userRouter := chi.NewRouter()

	userRouter.Use(StripTrailingSlash)
	userRouter.Use(DefaultHeaders)

	userRouter.Get("/", handlers.Repo.HandleGetUsers)
	userRouter.Post("/add", handlers.Repo.HandleAddUser)

	return userRouter
}
