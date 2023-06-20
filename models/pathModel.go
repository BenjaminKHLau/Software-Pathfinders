package models

import "gorm.io/gorm"

type Path struct {
	gorm.Model
	PathName        string
	PathDescription string
	Posts           []Post   `gorm:"foreignKey:PathID"`
	Cohorts         []Cohort `gorm:"foreignKey:PathID"`
	// CohortID        uint
}
