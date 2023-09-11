package main

import (
	"net/http"

	"BookingsWebsite/pkg/config"
	"BookingsWebsite/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// Create a MultiPlexer
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSruf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
