package main

import (
	"io"
	"net/http"
)

func ir(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Default route")
}

func dr(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog route")
}

func mr(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "John")
}

func main() {

	http.HandleFunc("/", ir)
	http.HandleFunc("/dog", dr)
	http.HandleFunc("/me/", mr)

	http.ListenAndServe(":8080", nil)

}
