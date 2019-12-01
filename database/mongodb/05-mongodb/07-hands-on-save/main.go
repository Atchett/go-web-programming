package main

import (
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/07-hands-on-save/controllers"
	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/07-hands-on-save/models"
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

func getSession() map[string]models.User {
	// will return data or empty map
	return models.LoadUsers()
}
