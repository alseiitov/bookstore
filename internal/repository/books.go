package repository

import (
	"database/sql"

	"github.com/alseiitov/bookstore/internal/model"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

func (r *BooksRepo) Create(book model.Book) error {
	_, err := r.db.Exec(`INSERT INTO books (name, author) VALUES ($1, $2)`, book.Name, book.Author)
	return err
}

func (r *BooksRepo) GetAll() ([]model.Book, error) {
	var books []model.Book

	rows, err := r.db.Query(`SELECT id, name, author FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Name, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BooksRepo) GetByID(bookID int) (model.Book, error) {
	var book model.Book

	err := r.db.QueryRow(
		`SELECT id, name, author FROM books WHERE id = $1`, bookID,
	).Scan(
		&book.ID, &book.Name, &book.Author,
	)

	return book, err
}

func (r *BooksRepo) Update(book model.Book) error {
	_, err := r.db.Exec(
		`UPDATE books SET name = $1, author = $2 WHERE id = $3`,
		book.Name, book.Author, book.ID,
	)

	return err
}

func (r *BooksRepo) Delete(bookID int) error {
	_, err := r.db.Exec(`DELETE FROM books WHERE id = $1`, bookID)
	return err
}
