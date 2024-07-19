package models

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string
}
