package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/01-update-controllers/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	// connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// check if connection error (is mongodb running)
	if err != nil {
		panic(err)
	}
	return s
}
