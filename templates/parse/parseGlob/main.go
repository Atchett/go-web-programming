package main

import (
	"text/template"
	"log"
	"os"
)

func main() {

	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "tpl.1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "tpl.2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
