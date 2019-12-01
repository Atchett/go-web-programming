package main

import (
	"io"
	"log"
	"net"
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

		// they all implement the type Writer
		io.WriteString(conn, "\nHello from TCP server\n")
		conn.Close()

	}
}
