package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	// read request
	request(conn)
}

func request(conn net.Conn) {

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// multiplexer - route what goes where
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	// get the 1st word of the 1st line
	// GETS the method
	m := strings.Fields(ln)[0]
	fmt.Println("**** METHOD", m)
	u := strings.Fields(ln)[1]
	fmt.Println("**** URL", u)

	if m == "GET" && u == "/" {
		index(conn)
	}

	if m == "GET" && u == "/about" {
		about(conn)
	}

	if m == "GET" && u == "/contact" {
		contact(conn)
	}

	if m == "GET" && u == "/apply" {
		apply(conn)
	}

	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html>
			<html land="en">
			<head>
				<meta charset="UTF-1">
				<title>Some test Doc</title>
			</head>
			<body>
				<strong>Index</strong></br>
				<a href="/">index</a></br>
				<a href="/about">about</a></br>
				<a href="/contact">contact</a></br>
				<a href="/apply">apply</a></br>
			</body>
		</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html>
			<html land="en">
			<head>
				<meta charset="UTF-1">
				<title>Some test Doc</title>
			</head>
			<body>
				<strong>About</strong></br>
				<a href="/">index</a></br>
				<a href="/about">about</a></br>
				<a href="/contact">contact</a></br>
				<a href="/apply">apply</a></br>
			</body>
		</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html>
			<html land="en">
			<head>
				<meta charset="UTF-1">
				<title>Some test Doc</title>
			</head>
			<body>
				<strong>Contact</strong></br>
				<a href="/">index</a></br>
				<a href="/about">about</a></br>
				<a href="/contact">contact</a></br>
				<a href="/apply">apply</a></br>
			</body>
		</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html>
			<html land="en">
			<head>
				<meta charset="UTF-1">
				<title>Some test Doc</title>
			</head>
			<body>
				<strong>Apply</strong></br>
				<a href="/">index</a></br>
				<a href="/about">about</a></br>
				<a href="/contact">contact</a></br>
				<a href="/apply">apply</a></br>
				<form method="POST" action="/apply">
					<input type="submit" value="apply">
				</form>
			</body>
		</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {

	body := `
		<!DOCTYPE html>
		<html>
			<html land="en">
			<head>
				<meta charset="UTF-1">
				<title>Some test Doc</title>
			</head>
			<body>
				<strong>Apply Process</strong></br>
				<a href="/">index</a></br>
				<a href="/about">about</a></br>
				<a href="/contact">contact</a></br>
				<a href="/apply">apply</a></br>
			</body>
		</html>
	`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
