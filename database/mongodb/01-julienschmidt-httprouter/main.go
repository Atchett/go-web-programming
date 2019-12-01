package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()
	r.GET("/", index)
	http.ListenAndServe(":8080", r)

}

// need the additional param, httprouter.Params (being thrown away with the _)
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
