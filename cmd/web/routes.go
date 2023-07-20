package main

import (
	"net/http"

	"github.com/kodega2016/booking-app/internal/config"
	"github.com/kodega2016/booking-app/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// make a reservation
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.SearchAvailabilty)
	mux.Post("/search-availability", handlers.Repo.PostSearchAvailabilty)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)
	return mux
}
