package models

import (
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
	ProfileImage string
	Verified     bool     `gorm:"default:false"`
	Approved     bool     `gorm:"default:false"`
	Parcels      []Parcel `gorm:"foreignKey:WorkerID"`
}

// 'CreatedAt' and 'UpdatedAt' automatically managed by GORM

type WorkerReqBody struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Contact      string `json:"contact"`
	ProfileImage string `json:"profile_image"`
	Verified     bool   `json:"verified"`
	Approved     bool   `json:"approved"`
}
