package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteConsoleNext logs the information about the request
func WriteConsoleNext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("requesting...", r.URL)
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds csrf protection for all the post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
