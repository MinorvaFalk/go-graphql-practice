package repository

import (
	"errors"
	"log"

	"github.com/MinorvaFalk/go-graphql-practice/core/utils"
	"github.com/MinorvaFalk/go-graphql-practice/models"
	"gorm.io/gorm"
)

type bookPostgresRepository struct {
	l  *log.Logger
	db *gorm.DB
}

func NewBookPostgresRepository(l *log.Logger, db *gorm.DB) BookRepository {
	return &bookPostgresRepository{l, db}
}

func (r *bookPostgresRepository) FetchAllBook() ([]models.Book, error) {
	var books []models.Book

	if err := r.db.Table("books").Find(&books).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorResultEmpty
		}

		r.l.Fatal(err)

		return nil, err
	}

	return books, nil
}

func (r *bookPostgresRepository) AddNewBook(book models.Book) (models.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookPostgresRepository) FetchBook(id string) (models.Book, error) {
	var book models.Book

	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		return book, err
	}

	return book, nil
}
