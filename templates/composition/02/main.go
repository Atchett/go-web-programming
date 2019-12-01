package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person{
			Name: "Ian Flemming",
			Age:  56,
		},
		false,
	}

	err := tmpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}

}
