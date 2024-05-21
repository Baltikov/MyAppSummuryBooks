package model

import (
	"fmt"
	"testapi/bd"
)

type Book struct {
	ID          int64  `json:"book_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Citation    string `json:"citation"`
}

func GetBooks() ([]Book, error) {
	query := `SELECT * FROM books`
	rowBooks, err := bd.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying books: %v", err)
	}
	defer rowBooks.Close()

	var books []Book

	for rowBooks.Next() {
		var book Book
		err := rowBooks.Scan(&book.ID, &book.Title, &book.Description, &book.Citation)
		if err != nil {
			return nil, fmt.Errorf("error scanning book: %v", err)
		}
		books = append(books, book)
	}

	if err = rowBooks.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over books: %v", err)
	}

	return books, nil
}
