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
	CreateTableBook()
	CreateTableFAQ()

}

func CreateTableBook() {
	bookTable := `CREATE TABLE IF NOT EXISTS books (
    book_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(55) NOT NULL,
    description TEXT NOT NULL,
    citation TEXT NOT NULL
);`

	_, err := DB.Exec(bookTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}
func CreateTableFAQ() {
	query := `
		CREATE TABLE IF NOT EXISTS faq (
			faq_id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT NOT NULL,
			category TEXT NOT NULL
		)
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	query = `
		INSERT INTO  faq (description, category)
		VALUES
			('как добавить книгу', 'Пользователь'),
			('как опубликовать книгу', 'Автор'),
			('как добавить Img для книги', 'Пользователь'),
			('как купить подписку', 'Пользователь'),
			('Куда писать если что-то не работает', 'Пользователь'),
			('test', 'Пользователь'),
			('test', 'Пользователь'),
			('Как монетезировать свою книгу в приложении', 'Автор'),
			('Куда писать если что-то не работает', 'Автор')
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
