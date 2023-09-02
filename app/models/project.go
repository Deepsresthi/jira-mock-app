package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	ProjectName  string `gorm:"not null"`
	TeamName     string `gorm:"not null"`
	ProjectOwner string `gorm:"-"`
	ProjectID    uint   `gorm:"not null"`
}

func (p *Project) BeforeCreate(scope *gorm.Scope) error {
	// Validate that ProjectOwner exists in the user database
	var user User
	if scope.DB().Where("username = ?", p.ProjectOwner).First(&user).RecordNotFound() {
		return gorm.ErrRecordNotFound
	}

	scope.SetColumn("CreatedAt", time.Now())
	return nil
}
