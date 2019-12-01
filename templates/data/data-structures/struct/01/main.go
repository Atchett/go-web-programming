package main

import (
	"text/template"
	"log"
	"os"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	buddah := sage{
		Name:  "Buddah",
		Motto: "The belief of no beliefs",
	}

	err := tmpl.Execute(os.Stdout, buddah)
	if err != nil {
		log.Fatalln(err)
	}
}
