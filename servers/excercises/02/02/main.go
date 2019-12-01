package main

import (
	"io"
	"log"
	"net"
	"bufio"
	"fmt"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		
		// go routine to handle the connection
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	// gets the line by line
	scanner := bufio.NewScanner(conn)
	// if there is a line more to the next token
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// they all implement the type Writer
	io.WriteString(conn, "I see you connected.")
	fmt.Println("Code got here.")
}
