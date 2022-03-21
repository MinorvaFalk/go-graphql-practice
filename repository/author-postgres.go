package repository

import (
	"errors"
	"log"

	"github.com/MinorvaFalk/go-graphql-practice/core/utils"
	"github.com/MinorvaFalk/go-graphql-practice/models"
	"gorm.io/gorm"
)

type authorPostgresRepository struct {
	l  *log.Logger
	db *gorm.DB
}

func NewAuthorPostgresRepository(l *log.Logger, db *utils.Database) AuthorRepository {
	return &authorPostgresRepository{l, db.DB}
}

func (r *authorPostgresRepository) FetchAllAuthor() ([]models.Author, error) {
	var authors []models.Author

	if err := r.db.Table("authors").Find(&authors).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorResultEmpty
		}

		r.l.Fatal(err)

		return nil, err
	}

	return authors, nil
}

func (r *authorPostgresRepository) AddNewAuthor(name string) (string, error) {
	if err := r.db.Create(&models.Author{Name: name}).Error; err != nil {
		return "", err
	}

	return name, nil
}

func (r *authorPostgresRepository) FetchAuthor(id string) (models.Author, error) {
	var author models.Author

	if err := r.db.First(&author, "id = ?", id).Error; err != nil {
		return author, err
	}

	return author, nil
}
