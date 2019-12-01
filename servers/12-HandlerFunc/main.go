package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "kit kit kitty")
}

func main() {
	// will catch /dog/something/else
	http.Handle("/dog/", http.HandlerFunc(d))
	// will not catch /cat/something/else
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
