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
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/del", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Successfully completed.")
}

func amigos(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)

	// data to be used
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)

}

func create(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)

}

func insert(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)

	defer stmt.Close()

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)

}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer`)
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Println(name)

		fmt.Fprintln(w, "RETREIVED RECORD:", name)
	}

}

func update(w http.ResponseWriter, r *http.Request) {

	stmt, err := db.Prepare(`UPDATE customer SET name = "Jimmy" WHERE name = "James";`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)

}

func del(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name = "Jimmy";`)
	check(err)

	res, err := stmt.Exec()
	check(err)

	n, err := res.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)

	_, err = stmt.Exec()
	check(err)

	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func check(err error) {

	if err != nil {
		fmt.Println(err)
	}

}
