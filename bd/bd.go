package bd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"

	"testapi/pkg/loger"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		loger.Logrus.Fatalf("Error opening DB: %v", err)
	} else {
		loger.Logrus.Println("Successfully opened DB")
		fmt.Println("Successfully opened DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	CreateTableUsers()
	CreateTableBook()
	CreateTableFAQ()

}

func CreateTableBook() {
	bookTable := `CREATE TABLE IF NOT EXISTS books (
    book_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(55) NOT NULL,
    description TEXT NOT NULL,
    citation TEXT NOT NULL,
    user_id INTEGER,
    FOREIGN KEY(user_id) REFERENCES users (id)
);`

	_, err := DB.Exec(bookTable)
	if err != nil {
		loger.Logrus.Fatal(err)
	}
	loger.Logrus.Info("Table created successfully")
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
func CreateTableUsers() {
	query := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email VARCHAR(255) UNIQUE NOT NULL,
        password TEXT NOT NULL
    )`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table Users created successfully")
	loger.Logrus.Info("Table Users created successfully")
}
