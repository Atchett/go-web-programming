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

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Spurgin",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
