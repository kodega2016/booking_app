package main

import (
	"fmt"
	"log"
	"net/http"
)

// port for the application
const port = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page")
}

// main is the main entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("web server is running on %s", port))
	http.ListenAndServe(port, nil)
}
