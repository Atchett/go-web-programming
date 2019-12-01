package main

import "fmt"

// output to file using
// go run main.go > index.html
func main() {
	name := "John Spurgin"

	tpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World</title>
		</head>
		<body>
		<h1>` + name + `</h1>
		</body>
		</html>
	`

	fmt.Println(tpl)
}
