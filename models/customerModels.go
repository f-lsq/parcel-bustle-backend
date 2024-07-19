package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string
	Contact   string `gorm:"unique;not null"`
}
