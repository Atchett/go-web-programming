package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

// make sure the value stored is the same as the
// value you get back

func main() {

	// test the hashing
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
	c = getCode("test@example.com")
	fmt.Println(c)

}

func getCode(s string) string {
	// creating a hash - using a secret key
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
