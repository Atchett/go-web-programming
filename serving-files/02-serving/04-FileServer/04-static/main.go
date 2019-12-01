package main

import (
	"log"
	"net/http"
)

func main() {
	// must have index.html to serve as static file server
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	// sensible option to do this
	// ListenAndServe returns an error this handles the error
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
