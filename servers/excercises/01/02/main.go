package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func ir(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func dr(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "dog.gohtml", nil)
	handleError(w, err)
}

func mr(w http.ResponseWriter, r *http.Request) {
	n := "John"
	err := tmpl.ExecuteTemplate(w, "me.gohtml", n)
	handleError(w, err)

}

func main() {

	http.HandleFunc("/", ir)
	http.HandleFunc("/dog", dr)
	http.HandleFunc("/me/", mr)

	http.ListenAndServe(":8080", nil)

}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
