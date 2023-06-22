package models

import (
	"time"

	"gorm.io/gorm"
)

type Cohort struct {
	gorm.Model
	StartDate time.Time
	PathID    uint
	Path      Path   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Belongs to Path
	Users     []User `gorm:"many2many:cohort_users"`                       // Users associated with the cohort
	// Path      Path // Belongs to Path
}
