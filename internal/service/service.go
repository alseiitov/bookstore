package service

import (
	"github.com/alseiitov/bookstore/internal/model"
	"github.com/alseiitov/bookstore/internal/repository"
)

type Books interface {
	Create(name, author string) error
	GetAll() ([]model.Book, error)
	GetByID(bookID int) (model.Book, error)
	Update(bookID int, name, author string) error
	Delete(bookID int) error
}

type Services struct {
	Books Books
}

type ServiceDeps struct {
	Repos *repository.Repositories
}

func NewServices(deps ServiceDeps) *Services {
	bookService := NewBooksService(deps.Repos.Books)
	return &Services{
		Books: bookService,
	}
}
