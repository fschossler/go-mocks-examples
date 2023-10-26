package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAll(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Define the expected query and result
	expectedRows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(1, "Book 1", "Author 1").
		AddRow(2, "Book 2", "Author 2").
		AddRow(3, "Book 3", "Author 3")

	mock.ExpectQuery("SELECT \\* FROM books").WillReturnRows(expectedRows)

	// Call the function under test
	rows, err := getAll(db)
	if err != nil {
		t.Fatalf("Error in getAll: %v", err)
	}
	defer rows.Close()

	// Iterate through the result rows and print the data
	for rows.Next() {
		var id int
		var title, author string
		if err := rows.Scan(&id, &title, &author); err != nil {
			t.Fatalf("Error scanning rows: %v", err)
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", id, title, author)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("Error in rows: %v", err)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
