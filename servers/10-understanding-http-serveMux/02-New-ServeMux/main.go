package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kit kit kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	// will catch /dog/something/else
	mux.Handle("/dog/", d)
	// will not catch /cat/something/else
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
