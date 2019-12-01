package main

import (
	"text/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	sages := []string{
		"Ghandi",
		"MLK",
		"Buddha",
		"Jesus",
		"Muhammad",
	}

	err := tmpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
