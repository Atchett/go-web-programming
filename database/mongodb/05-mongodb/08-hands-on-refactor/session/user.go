package session

import (
	"encoding/json"
	"fmt"
	"os"

	"bitbucket.org/johnpersonal/go-web-programming/mongodb/05-mongodb/08-hands-on-refactor/models"
)

// GetSession loads the user data or returns empty map
// for storing session data
func GetSession() map[string]models.User {
	// will return data or empty map
	return LoadUsers()
}

// StoreUsers stores the user map
func StoreUsers(m map[string]models.User) {

	// create a file
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// write to the file (f) with encoded map (m)
	json.NewEncoder(f).Encode(m)
}

// LoadUsers loads the user map on startup
func LoadUsers() map[string]models.User {
	m := make(map[string]models.User)

	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		// return empty map
		return m
	}
	defer f.Close()

	// decode file into the map
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}
