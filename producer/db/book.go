package db

import (
	"database/sql"

	"github.com/ornellast/bookstore/producer/models"
)

func (db Database) GetAllBooks() (*models.BookList, error) {
	list := &models.BookList{}
	rows, err := db.Conn.Query("SELECT * FROM books ORDER BY id DESC")

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Description, &book.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Books = append(list.Books, book)
	}
	return list, nil
}

// func (db Database) AddBook(book *models.Book) error {
// 	var id int
// 	var createdAt string
// 	query := `INSERT INTO books (name, description) VALUES($1, $2) RETURNING id, created_at`
// 	err := db.Conn.QueryRow(query, book.Name, book.Description).Scan(&id, &createdAt)

// 	if err != nil {
// 		return err
// 	}
// 	book.ID = id
// 	book.CreatedAt = createdAt
// 	return nil
// }

func (db Database) GetBookById(bookId int) (models.Book, error) {
	book := models.Book{}
	qy := `SELECT * FROM books WHERE id=$1`
	row := db.Conn.QueryRow(qy, bookId)
	switch err := row.Scan(&book.ID, &book.Name, &book.Description, &book.CreatedAt); err {
	case sql.ErrNoRows:
		return book, ErrNoMatch
	default:
		return book, err
	}
}

// func (db Database) DeleteBook(bookId int) error {
// 	qy := `DELETE FROM books WHERE id = $1`
// 	_, err := db.Conn.Exec(qy, bookId)

// 	switch err {
// 	case sql.ErrNoRows:
// 		return ErrNoMatch
// 	default:
// 		return err
// 	}
// }

// func (db Database) UpdateBook(bookId int, bookData models.Book) (models.Book, error) {
// 	book := models.Book{}
// 	query := `UPDATE books SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`

// 	err := db.Conn.QueryRow(query, bookData.Name, bookData.Description, bookId).Scan(&book.ID, &book.Name, &book.Description, &book.CreatedAt)

// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			return book, ErrNoMatch
// 		}
// 		return book, err
// 	}
// 	return book, nil
// }
