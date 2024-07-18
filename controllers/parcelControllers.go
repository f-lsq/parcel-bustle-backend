package controllers

import (
	"fmt"
	"time"

	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/gin-gonic/gin"
)

func CreateParcel(c *gin.Context) {
	var body struct {
		Status          string `json:"status"`
		DeliveryAddress string `json:"delivery_address"`
		ReturnAddress   string `json:"return_address"`
		DeliverBy       string `json:"deliver_by"`
		DeliveredImage  string `json:"delivered_image"`
		PaymentType     string `json:"payment_type"`
		PaymentMade     bool   `json:"payment_made"`
	}

	c.BindJSON(&body)

	var DeliverByDate time.Time
	if body.DeliverBy == "" {
		DeliverByDate = time.Now().AddDate(0, 0, 7)
	} else {
		DeliverByDate, _ = time.Parse("2006-01-02", body.DeliverBy)
	}

	parcel := models.Parcel{
		Status:          body.Status,
		DeliveryAddress: body.DeliveryAddress,
		ReturnAddress:   body.ReturnAddress,
		DeliverBy:       DeliverByDate,
		DeliveredImage:  body.DeliveredImage,
		PaymentType:     body.PaymentType,
		PaymentMade:     body.PaymentMade,
	}

	fmt.Println(parcel)
	result := config.DB.Create(&parcel)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parcel created successfully!",
		"data":    parcel,
	})

}
