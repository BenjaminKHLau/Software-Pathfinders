package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"unique"`
	Posts     []Post //`gorm:"foreignKey:AuthorID"`
	Admin     bool
}

// Cohort    Cohort
// CohortID  int
