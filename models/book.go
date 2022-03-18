package models

import (
	"time"
)

type Book struct {
	ID          uint `json:"id"`
	Title       string
	Description string
	Rating      float64
	AuthorRefer uint      `gqlgen:"authorId" json:"authorId"`
	Author      Author    `gorm:"foreignKey:AuthorRefer" json:"-" gqlgen:"-"`
	CreateAt    time.Time `json:"-"`
	UpdateAt    time.Time `json:"-"`
}
