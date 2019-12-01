package controllers

import (
	"html/template"
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/09-hands-on-refactor-session-expire/session"
)

// Controller to handle the controllers
type Controller struct {
	tpl *template.Template
}

// NewController creates a new controller struct
// to access methods in controller
func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

// Index handles the default route
func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// Bar handles the Bar route
func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
