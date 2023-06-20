package models

import (
	"time"

	"gorm.io/gorm"
)

type Cohort struct {
	gorm.Model
	StartDate time.Time
	UserID    []int
	PathID    int
	Path      Path `gorm:"foreignKey:PathID"`
}
