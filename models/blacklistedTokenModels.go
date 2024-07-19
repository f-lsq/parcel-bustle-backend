package models

import "gorm.io/gorm"

type BlacklistedToken struct {
	gorm.Model
	FirstName string `gorm:"not null"`
}
