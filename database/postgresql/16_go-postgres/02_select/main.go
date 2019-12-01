package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Book struct for the bookstore db row data
type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		// scan the row data into the pointer to the Book struct
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // order matters
		if err != nil {
			panic(err)
		}
		// add the Book struct to the Books slice
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	// iterate over the data
	for _, bk := range bks {
		fmt.Printf("%s, %s, %s, %.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}

}
