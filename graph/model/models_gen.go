// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type BookInput struct {
	Title       string   `json:"title"`
	Description *string  `json:"description"`
	Rating      *float64 `json:"rating"`
	AuthorID    int      `json:"authorId"`
}