package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

type user struct {
	Username string
	Password []byte
	First    string
	Last     string
	Role     string
}

var dbUsers = map[string]user{}
var dbSessions = map[string]string{}
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
	// bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	// dbUsers["test@test.com"] = user{
	// 	Username: "test@test.com",
	// 	Password: bs,
	// 	First:    "John",
	// 	Last:     "Test",
	// }
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be a 007 to enter the bar", http.StatusForbidden)
		return
	}
	tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		rl := r.FormValue("role")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session ID
		sID, _ := uuid.NewV4()

		// cooke
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		// store the session
		dbSessions[c.Value] = un

		// store in users
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Error generating password hash", http.StatusInternalServerError)
		}
		u := user{
			Username: un,
			Password: bs,
			First:    f,
			Last:     l,
			Role:     rl,
		}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)

}

func login(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process the form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// is there a username
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username / Password don't match", http.StatusForbidden)
			return
		}

		// does the encrypted password match stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username / Password don't match", http.StatusForbidden)
			return
		}

		// create session cookie
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		// store in sessions store
		dbSessions[c.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tmpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")

	// remove entry from session map
	delete(dbSessions, c.Value)

	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// redirect user
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
