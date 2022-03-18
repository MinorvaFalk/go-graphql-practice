package services

import (
	gqlmodel "github.com/MinorvaFalk/go-graphql-practice/graph/model"
	"github.com/MinorvaFalk/go-graphql-practice/models"
)

type BookService interface {
	FetchAllBook() ([]*models.Book, error)
	AddNewBook(input gqlmodel.BookInput) (*models.Book, error)
	FetchBook(id string) (*models.Book, error)
}

type AuthorService interface {
	FetchAllAuthor() ([]*models.Author, error)
	AddNewAuthor(name string) (string, error)
	FetchAuthor(id string) (*models.Author, error)
}
