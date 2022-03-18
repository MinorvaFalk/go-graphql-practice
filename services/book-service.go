package services

import (
	"log"

	gqlmodel "github.com/MinorvaFalk/go-graphql-practice/graph/model"
	"github.com/MinorvaFalk/go-graphql-practice/models"
	"github.com/MinorvaFalk/go-graphql-practice/repository"
)

type bookService struct {
	l    *log.Logger
	repo repository.BookRepository
}

func NewBookService(l *log.Logger, repo repository.BookRepository) BookService {
	return &bookService{l, repo}
}

func (s *bookService) FetchAllBook() ([]*models.Book, error) {
	var books []*models.Book

	result, err := s.repo.FetchAllBook()
	if err != nil {
		return nil, err
	}

	for idx := range result {
		books = append(books, &result[idx])
	}

	return books, nil
}

func (s *bookService) AddNewBook(input gqlmodel.BookInput) (*models.Book, error) {
	book := models.Book{
		Title:       input.Title,
		Description: *input.Description,
		Rating:      *input.Rating,
		AuthorRefer: uint(input.AuthorID),
	}

	result, err := s.repo.AddNewBook(book)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *bookService) FetchBook(id string) (*models.Book, error) {
	result, err := s.repo.FetchBook(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
