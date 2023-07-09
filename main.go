package main

import (
	"fmt"
	"log"
	"net/http"
)

// port for the application
const port = ":8080"

// main is the main entry point of the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("web server is running on %s", port))
	http.ListenAndServe(port, nil)
}
