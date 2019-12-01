package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// User models the user
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	ID     string `json:"id"`
}

// StoreUsers stores the user map
func StoreUsers(m map[string]User) {

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
func LoadUsers() map[string]User {
	m := make(map[string]User)

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
