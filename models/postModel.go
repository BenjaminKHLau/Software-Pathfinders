package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string
	Body   string
	User   User // Belongs to User
	UserID uint
	Path   Path // Belongs to Path
	PathID uint
}
