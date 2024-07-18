package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Contact      string `gorm:"not null"`
	ProfileImage sql.NullString
	Verified     bool `gorm:"default:false"`
	Approved     bool `gorm:"default:false"`
}

// 'CreatedAt' and 'UpdatedAt' automatically managed by GORM
