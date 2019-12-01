package models

import "time"

// User defines the User struct
type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

// Session defines the Session struct
type Session struct {
	UserName     string
	LastActivity time.Time
}
