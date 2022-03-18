package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	gqlmodel "github.com/MinorvaFalk/go-graphql-practice/graph/model"
	"github.com/MinorvaFalk/go-graphql-practice/models"
)

func (r *authorResolver) ID(ctx context.Context, obj *models.Author) (int, error) {
	return int(obj.ID), nil
}

func (r *bookResolver) ID(ctx context.Context, obj *models.Book) (int, error) {
	return int(obj.ID), nil
}

func (r *bookResolver) AuthorID(ctx context.Context, obj *models.Book) (string, error) {
	return fmt.Sprint(obj.AuthorRefer), nil
}

func (r *mutationResolver) AddAuthor(ctx context.Context, name string) (string, error) {
	return r.As.AddNewAuthor(name)
}

func (r *mutationResolver) AddBook(ctx context.Context, book gqlmodel.BookInput) (*models.Book, error) {
	return r.Bs.AddNewBook(book)
}

func (r *queryResolver) Status(ctx context.Context) (string, error) {
	return "running", nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
	return r.Bs.FetchAllBook()
}

func (r *queryResolver) Book(ctx context.Context, id string) (*models.Book, error) {
	return r.Bs.FetchBook(id)
}

func (r *queryResolver) Authors(ctx context.Context) ([]*models.Author, error) {
	return r.As.FetchAllAuthor()
}

func (r *queryResolver) Author(ctx context.Context, id string) (*models.Author, error) {
	return r.As.FetchAuthor(id)
}

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
