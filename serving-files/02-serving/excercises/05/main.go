package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("starting-files/templates/*"))
}

func main() {

	fs := http.StripPrefix("/public", http.FileServer(http.Dir("./starting-files/public")))
	http.Handle("/public/", fs)
	http.HandleFunc("/", iRt)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func iRt(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
