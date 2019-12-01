package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// setup the template container and parse a file into it
	tmpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute the template
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// add more template files into the container
	tmpl, err = tmpl.ParseFiles("tpl.1.gohtml", "tpl.2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// execute the named template
	err = tmpl.ExecuteTemplate(os.Stdout, "tpl.1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute the named template
	err = tmpl.ExecuteTemplate(os.Stdout, "tpl.2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// executes the first template found in the container
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
