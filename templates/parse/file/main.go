package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// parse the file into the container for templates
	tmpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// create a new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	// defer the closing of the file to the end of the function
	defer nf.Close()

	// write the template to the new file
	err = tmpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
