package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getAll(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM books")
	return rows, err
}

func main() {
	db, err := sql.Open("mysql", "myuser:userpassword@tcp(localhost:3306)/books")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := getAll(db)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the result rows and print the data
	for rows.Next() {
		var id int
		var title, author string
		if err := rows.Scan(&id, &title, &author); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", id, title, author)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
