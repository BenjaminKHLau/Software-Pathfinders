package models

import (
	"time"

	"gorm.io/gorm"
)

type Cohort struct {
	gorm.Model
	StartDate time.Time
	// UserID    []int
	PathID uint
	Path   Path `gorm:"foreignKey:PathID"`
	// CohortID uint
}
