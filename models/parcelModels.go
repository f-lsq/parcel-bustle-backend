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

type ParcelReqBody struct {
	Status          string `json:"status"`
	DeliveryAddress string `json:"delivery_address"`
	ReturnAddress   string `json:"return_address"`
	DeliverBy       string `json:"deliver_by"`
	DeliveredImage  string `json:"delivered_image"`
	PaymentType     string `json:"payment_type"`
	PaymentMade     bool   `json:"payment_made"`
}
