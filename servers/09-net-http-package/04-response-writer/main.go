package main

import (
	"fmt"
	"net/http"
)

// just silly name to show type can be called anything
type hotdog int

// method on the type
// hotdog now a value of type Handler
// implicityly implements the Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Johns-Key", "This is my key")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "Any code you want in here")

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
