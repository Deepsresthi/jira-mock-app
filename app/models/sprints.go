package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Sprints struct {
	gorm.Model
	SprintID  uint `gorm:"primaryKey"`
	Name      string
	StartDate time.Time
	EndDate   time.Time
}
