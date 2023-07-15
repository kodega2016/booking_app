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

	// use load session load
	mux.Use(SessionLoad)

	// file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/about", handlers.Repo.About)
	mux.HandleFunc("/make-reservation", handlers.Repo.Reservation)
	mux.HandleFunc("/generals-quarters", handlers.Repo.Generals)
	mux.HandleFunc("/majors-suite", handlers.Repo.Majors)
	mux.HandleFunc("/search-availability", handlers.Repo.SearchAvailabilty)
	mux.HandleFunc("/contact", handlers.Repo.Contact)
	return mux
}
