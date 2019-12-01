package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tmpl *template.Template

// composite literal - different way for map to be made
var dbUsers = map[string]user{}      // userID, user
var dbSessions = map[string]string{} // session ID, user ID
// or different way to make a map
//var dbSessions = make(map[string]string)

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	//get the cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID, err := uuid.NewV4()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get the user
	var u user
	// comma ok idiom
	// if value not there will give back the zero value and false
	// if value is there will give back value and true
	// this is getting value (un, ok := dbSessions[c.Value]) then checking true (; ok)
	// if true assign the user to the u var
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process the form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{
			UserName: un,
			First:    f,
			Last:     l,
		}
		// store the user id (username in this case)
		// with the key of session id (cookie value)
		dbSessions[c.Value] = un
		// store the user with the key of username
		dbUsers[un] = u
	}

	// execute the template
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}
