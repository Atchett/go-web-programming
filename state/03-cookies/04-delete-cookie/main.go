package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "some-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read cookie</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("some-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		// return out of the function after a redirect
		return
	}
	fmt.Fprintf(w, `<h1>Your cookie: <br> %v</h1>
					<h2><a href="/expire">Expire cookie</a></h2>`, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("some-cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		// return out of the function after a redirect
		return
	}
	c.MaxAge = -1 // delete the cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
