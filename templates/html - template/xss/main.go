package main

import (
	"log"
	"os"
	"text/template"
)

// Page is example
type Page struct {
	Title   string
	Heading string
	Input   string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	home := Page{
		Title:   "Nothing is escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Yow!")</script>`,
	}

	err := tmpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}

}
