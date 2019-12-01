package main

import (
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Foo ran")
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}
