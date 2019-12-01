package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"upCase":     strings.ToUpper,
	"fThreeChar": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	// have to add the funcs to a new / empty template container before
	// you can parse the files into the container
	// otherwise the template doesn't have a reference to the funcs
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {

	b := sage{
		Name:  "Buddah",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	bm := car{
		Manufacturer: "BMW",
		Model:        "435i",
		Doors:        4,
	}

	sages := []sage{
		b,
		g,
		m,
	}

	cars := []car{
		f,
		bm,
	}

	// inline defining the struct
	// anonymous type
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	// execute the named template
	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
