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
	UserID      int64  `json:"user_id"`
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
func GetBook(id int64) (Book, error) {
	query := `SELECT * FROM books WHERE id = userId`
	rowBook, err := bd.DB.Query(query, id)
	if err != nil {
		return Book{}, fmt.Errorf("error querying book: %v", err)
	}
	var book Book
	err = rowBook.Scan(&book.ID, &book.Title, &book.Description, &book.Citation)
	if err != nil {
		return Book{}, fmt.Errorf("error scanning book: %v", err)
	}
	return book, nil
}
func CreateBook(book Book) error {
	query := "INSERT INTO books(title, description, citation, user_id) VALUES(?, ?, ?, ?)"
	result, err := bd.DB.Exec(query, &book.Title, &book.Description, &book.Citation, &book.UserID)
	if err != nil {
		return fmt.Errorf("error inserting book: %v", err)
	}
	book.ID, err = result.LastInsertId()

	if err != nil {
		return fmt.Errorf("error getting books: %v", err)
	}
	fmt.Println()
	return nil

}

func UpdateBook(book Book) ([]Book, error) {
	query := `UPDATE books SET title = ?, description = ?, citation = ? WHERE book_id = ?`
	bd.DB.Exec(query, &book.Title, &book.Description, &book.Citation)
	books, err := GetBooks()
	if err != nil {
		return nil, fmt.Errorf("error getting books: %v", err)
	}
	return books, nil
}

func DeleteBook(id int64) error {
	query := `DELETE FROM books  WHERE book_id = ?`
	bd.DB.Exec(query, id)
	return nil
}
