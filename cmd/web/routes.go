package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/kodega2016/booking-app/pkg/config"
	"github.com/kodega2016/booking-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.HandleFunc("/", handlers.Repo.Home)
	mux.HandleFunc("/about", handlers.Repo.About)

	return mux
}
