package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/06-hands-on/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

// UserController models the usercontroller
type UserController struct {
	session map[string]models.User
}

// NewUserController creates a pointer to a UserController struct
func NewUserController(s map[string]models.User) *UserController {
	return &UserController{s}
}

// GetUser gets the user
// Methods have to be capitalised to be exported e.g. GetUser not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// grab ID
	id := p.ByName("id")

	// Fetch the user
	u := uc.session[id]

	// marshall user into JSON
	// throwing away the error
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

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

	// create ID
	id, _ := uuid.NewV4()

	u.ID = id.String()

	// store the user
	uc.session[u.ID] = u

	// marshall / unmarshall for having JSON assigned to a variable
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

}

// DeleteUser deletes the user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", id, "\n")

}
