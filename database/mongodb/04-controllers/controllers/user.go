package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/04-controllers/models"
	"github.com/julienschmidt/httprouter"
)

// UserController models the usercontroller
type UserController struct{}

// NewUserController creates a pointer to a UserController struct
func NewUserController() *UserController {
	return &UserController{}
}

// GetUser gets the user
// Methods have to be capitalised to be exported e.g. GetUser not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

// CreateUser creates the user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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

// DeleteUser deletes the user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")

}
