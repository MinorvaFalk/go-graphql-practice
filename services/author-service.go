package services

import (
	"log"

	"github.com/MinorvaFalk/go-graphql-practice/models"
	"github.com/MinorvaFalk/go-graphql-practice/repository"
)

type authorService struct {
	l    *log.Logger
	repo repository.AuthorRepository
}

func NewAuthorService(l *log.Logger, repo repository.AuthorRepository) AuthorService {
	return &authorService{l, repo}
}

func (s *authorService) FetchAllAuthor() ([]*models.Author, error) {
	var authors []*models.Author

	result, err := s.repo.FetchAllAuthor()
	if err != nil {
		return nil, err
	}

	for idx := range result {
		authors = append(authors, &result[idx])
	}

	return authors, nil
}

func (s *authorService) AddNewAuthor(name string) (string, error) {
	result, err := s.repo.AddNewAuthor(name)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *authorService) FetchAuthor(id string) (*models.Author, error) {
	result, err := s.repo.FetchAuthor(id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
