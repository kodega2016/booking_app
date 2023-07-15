package main

import (
	"fmt"
	"github.com/kodega2016/booking-app/internal/config"
	"github.com/kodega2016/booking-app/internal/handlers"
	"github.com/kodega2016/booking-app/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// port for the application
const port = ":8080"

// app for app config
var app config.AppConfig

// session for session manager
var session *scs.SessionManager

// main is the main entry point of the application
func main() {
	// change this to true when in production
	app.InProduction = false

	// setup session for the app
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// set session in app config
	app.Session = session
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Unable to create template cache.")
	}

	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	log.Println(fmt.Sprintf("web server is running on %s", port))
	// http.ListenAndServe(port, nil)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
