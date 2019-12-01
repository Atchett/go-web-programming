package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"fdateDYM": monthDayYear,
}

func init() {

	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}

func monthDayYear(t time.Time) string {
	// formatting time :
	// day = 01
	// month = 02
	// year = 06
	return t.Format("02-01-2006")
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
