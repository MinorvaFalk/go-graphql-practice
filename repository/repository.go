package repository

import "github.com/MinorvaFalk/go-graphql-practice/models"

type BookRepository interface {
	FetchAllBook() ([]models.Book, error)
	AddNewBook(models.Book) (models.Book, error)
	FetchBook(id string) (models.Book, error)
}

type AuthorRepository interface {
	FetchAllAuthor() ([]models.Author, error)
	AddNewAuthor(name string) (string, error)
	FetchAuthor(id string) (models.Author, error)
}
