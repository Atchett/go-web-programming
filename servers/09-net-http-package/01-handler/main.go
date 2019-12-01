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
	fmt.Println("Any code in here")
}

func main() {

}
