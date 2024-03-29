package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// just silly name to show type can be called anything
type hotdog int

// method on the type
// hotdog now a value of type Handler
// implicityly implements the Handler
// this is clunky!
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
	}{
		r.Method,
		r.URL,
		r.Form,
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	var d hotdog
	http.ListenAndServe(":8080", d)

}
