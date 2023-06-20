package models

import "gorm.io/gorm"

type Path struct {
	gorm.Model
	PathName        string
	PathDescription string
	Posts           []Post
	Cohorts         []Cohort
}
