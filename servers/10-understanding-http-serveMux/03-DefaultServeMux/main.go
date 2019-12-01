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

	// will catch /dog/something/else
	http.Handle("/dog/", d)
	// will not catch /cat/something/else
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
