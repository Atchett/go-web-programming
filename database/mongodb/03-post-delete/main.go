package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/02-json/models"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// added route
	r.POST("/user", createUser)
	// added route plus paramter
	r.DELETE("/user/:id", deleteUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	s := `<!DOCTYPE html>
		<html>
		<head>
		<meta charset="utf-8">
		<title>Index</title>
		</head>
		<body>
		<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
		</body>
		</html>`

	w.Header().Set("Content-Type", "text/html; charset=UTF-8;")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))

}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	// marshall user into JSON
	// throwing away the error
	uj, _ := json.Marshal(u)

	// write header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// composite literal - type and curly braces
	u := models.User{}

	// encode / decode for sending / receiving JSON tp / from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// change the ID
	u.ID = "007"

	// marshall / unmarshall for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// Write content type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")

}
