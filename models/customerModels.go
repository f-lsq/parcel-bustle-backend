package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string
	Contact   string   `gorm:"unique;not null"`
	Parcels   []Parcel `gorm:"foreignKey:CustomerID"`
}

// 'CreatedAt' and 'UpdatedAt' automatically managed by GORM

type CustomerReqBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Contact   string `json:"contact"`
}
