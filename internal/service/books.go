package service

import (
	"github.com/alseiitov/bookstore/internal/model"
	"github.com/alseiitov/bookstore/internal/repository"
)

type BookService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(name, author string) error {
	return s.repo.Create(
		model.Book{
			Name:   name,
			Author: author,
		},
	)
}

func (s *BookService) GetAll() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetByID(bookID int) (model.Book, error) {
	book, err := s.repo.GetByID(bookID)

	if err == repository.ErrBookNotFound {
		return book, ErrBookNotFound
	}

	return book, err
}

func (s *BookService) Update(bookID int, name, author string) error {
	err := s.repo.Update(
		model.Book{
			ID:     bookID,
			Name:   name,
			Author: author,
		},
	)

	if err == repository.ErrBookNotFound {
		return ErrBookNotFound
	}

	return err
}

func (s *BookService) Delete(bookID int) error {
	return s.repo.Delete(bookID)
}
