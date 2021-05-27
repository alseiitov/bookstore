package repository

import (
	"database/sql"

	"github.com/alseiitov/bookstore/internal/model"
)

type Books interface {
	Create(book model.Book) error
	GetAll() ([]model.Book, error)
	GetByID(bookID int) (model.Book, error)
	Update(book model.Book) error
	Delete(bookID int) error
}

type Repositories struct {
	Books Books
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Books: NewBooksRepo(db),
	}
}
