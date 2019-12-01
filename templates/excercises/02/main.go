package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Name   string
	Hotels []hotel
}

type regions []region

func init() {
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h := regions{
		region{
			Name: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "A hotel",
					Address: "AAA",
					City:    "AA",
					Zip:     "AA-AA",
				},
				hotel{
					Name:    "B hotel",
					Address: "BBB",
					City:    "BB",
					Zip:     "BB-BB",
				},
				hotel{
					Name:    "C hotel",
					Address: "CCC",
					City:    "CC",
					Zip:     "CC-CC",
				},
			},
		},
		region{
			Name: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "A hotel",
					Address: "AAA",
					City:    "AA",
					Zip:     "AA-AA",
				},
				hotel{
					Name:    "B hotel",
					Address: "BBB",
					City:    "BB",
					Zip:     "BB-BB",
				},
				hotel{
					Name:    "C hotel",
					Address: "CCC",
					City:    "CC",
					Zip:     "CC-CC",
				},
			},
		},
		region{
			Name: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "A hotel",
					Address: "AAA",
					City:    "AA",
					Zip:     "AA-AA",
				},
				hotel{
					Name:    "B hotel",
					Address: "BBB",
					City:    "BB",
					Zip:     "BB-BB",
				},
				hotel{
					Name:    "C hotel",
					Address: "CCC",
					City:    "CC",
					Zip:     "CC-CC",
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}

}
