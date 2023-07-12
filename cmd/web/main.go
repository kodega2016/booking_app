package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kodega2016/booking-app/pkg/config"
	"github.com/kodega2016/booking-app/pkg/handlers"
	"github.com/kodega2016/booking-app/pkg/render"
)

// port for the application
const port = ":8080"

// main is the main entry point of the application
func main() {
	var app config.AppConfig
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
