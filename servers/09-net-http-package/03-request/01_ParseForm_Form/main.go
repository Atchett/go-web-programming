package main

import (
	"html/template"
	"log"
	"net/http"
)

// just silly name to show type can be called anything
type hotdog int

// method on the type
// hotdog now a value of type Handler
// implicityly implements the Handler
// this is clunky!
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tmpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	var d hotdog
	http.ListenAndServe(":8080", d)

}
