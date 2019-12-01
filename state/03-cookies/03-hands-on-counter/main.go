package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("counter")

	// cookie not found
	if err == http.ErrNoCookie {
		// create the cookie
		cookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	// convert the cookie value to int
	// Atoi = Ascii to int
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatal(err)
	}
	// increment the count
	count++
	// assign the value to the cookie
	// conver to a string
	// Itoa = Int to ascii
	cookie.Value = strconv.Itoa(count)
	// set the cookie
	http.SetCookie(w, cookie)

	// Execut the template
	err = tmpl.ExecuteTemplate(w, "index.gohtml", cookie.Value)
	if err != nil {
		fmt.Println(err)
	}
}
