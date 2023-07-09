package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kodega2016/booking-app/pkg/handlers"
)

// port for the application
const port = ":8080"

// main is the main entry point of the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("web server is running on %s", port))
	http.ListenAndServe(port, nil)
}
