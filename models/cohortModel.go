package models

import (
	"time"

	"gorm.io/gorm"
)

type Cohort struct {
	gorm.Model
	StartDate time.Time
	PathID    uint
	Path      Path // Belongs to Path
}
