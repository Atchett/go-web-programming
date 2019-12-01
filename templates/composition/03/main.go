package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{
					Number: "CSCI-40",
					Name:   "Introduction to programming in Go",
					Units:  "4",
				},
				course{
					Number: "CSCI-130",
					Name:   "Introduction to Web programming with Go",
					Units:  "4",
				},
				course{
					Number: "CSCI-140",
					Name:   "Mobile Apps using Go",
					Units:  "4",
				},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{
					Number: "CSCI-50",
					Name:   "Advanced Go",
					Units:  "5",
				},
				course{
					Number: "CSCI-190",
					Name:   "Advanced Web programming with Go",
					Units:  "5",
				},
				course{
					Number: "CSCI-191",
					Name:   "Advanced Mobile Apps with Go",
					Units:  "5",
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}

}
