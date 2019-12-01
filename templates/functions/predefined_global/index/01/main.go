package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	xs := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}

}
