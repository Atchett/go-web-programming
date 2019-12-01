package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/04-update-user-controller-delete/models"
	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

// UserController models the usercontroller
type UserController struct {
	session *mgo.Session
}

// NewUserController creates a pointer to a UserController struct
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser gets the user
// Methods have to be capitalised to be exported e.g. GetUser not getUser
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// grab ID
	id := p.ByName("id")

	// verify id is ObjectId hex representation,
	// otherwise status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIDHex returns an ObjectId from
	// the provided hex representation
	oid := bson.ObjectIdHex(id)

	// composite literal of user
	u := models.User{}

	// Fetch the user from mongodb
	if err := uc.session.DB("go-web-deb-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		return
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
	u.ID = bson.NewObjectId()

	// store the user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

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
