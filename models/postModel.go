package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Body     string
	Author   User `gorm:"foreignKey:AuthorID"`
	AuthorID uint
}
