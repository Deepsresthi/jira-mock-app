package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"not null"`
	MailID       string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`
}


