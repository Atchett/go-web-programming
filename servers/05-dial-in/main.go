package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// read from the connection
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}
