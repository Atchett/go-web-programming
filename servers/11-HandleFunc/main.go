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
	http.HandleFunc("/dog/", d)
	// will not catch /cat/something/else
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
