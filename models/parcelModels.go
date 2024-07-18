package models

import (
	"time"

	"gorm.io/gorm"
)

type Parcel struct {
	gorm.Model
	Status          string    `gorm:"not null"`
	DeliveryAddress string    `gorm:"not null"`
	ReturnAddress   string    `gorm:"not null"`
	DeliverBy       time.Time `gorm:"not null;"`
	DeliveredImage  string
	PaymentType     string `gorm:"not null"`
	PaymentMade     bool   `gorm:"not null"`
}

// 'CreatedAt' and 'UpdatedAt' automatically managed by GORM
