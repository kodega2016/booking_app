package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kodega2016/booking-app/pkg/config"
	"github.com/kodega2016/booking-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// logger middleware
	mux.Use(middleware.Logger)

	// use custom middlware
	mux.Use(WriteConsoleNext)

	// use no surf
	mux.Use(NoSurf)

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/about", handlers.Repo.About)
	return mux
}
