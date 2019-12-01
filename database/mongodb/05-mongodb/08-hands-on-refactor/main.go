package main

import (
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/08-hands-on-refactor/session"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/08-hands-on-refactor/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(session.GetSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", r)
}
