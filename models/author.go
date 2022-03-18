package models

import "time"

type Author struct {
	ID       uint `json:"id"`
	Name     string
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
