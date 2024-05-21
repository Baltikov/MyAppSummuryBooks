package bd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	} else {
		log.Println("Successfully opened DB")
		fmt.Println("Successfully opened DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

}

func CreateTableBook() {
	bookTable := `CREATE TABLE IF NOT EXISTS book (
    book_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(55) NOT NULL,
    description TEXT NOT NULL,
    citate TEXT NOT NULL
);`

	_, err := DB.Exec(bookTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}
