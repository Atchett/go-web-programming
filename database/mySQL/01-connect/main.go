package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	// imported but not used the
	// _ indicates that it's an alias / thrown away
	// only used for setup
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Oscar123@tcp(localhost:3306)/testdb01?charset=utf8")
	check(err)

	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
}

func check(err error) {

	if err != nil {
		fmt.Println(err)
	}

}
