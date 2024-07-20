package controllers

import (
	"errors"
	"fmt"

	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateParcel(c *gin.Context) {
	var body models.ParcelReqBody

	// parse JSON data in req body and binds it to the ParcelReqBody struct
	// similar to req.body with app.use(express.json())
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	parcel, err := services.CreateParcel(&body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parcel created successfully!",
		"data":    parcel,
	})
}

func GetParcels(c *gin.Context) {
	parcels, err := services.GetAllParcels()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "All parcels retrieved successfully!",
		"data":    parcels,
	})
}

func GetParcelByID(c *gin.Context) {
	id := c.Param("id")
	parcel, err := services.GetParcelByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No parcel of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parcel retrieved successfully!",
		"data":    parcel,
	})
}

func UpdateParcel(c *gin.Context) {
	var body models.ParcelReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	parcel, err := services.UpdateParcelByID(id, body)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No parcel of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parcel updated successfully!",
		"data":    parcel,
	})
}

func DeleteParcel(c *gin.Context) {
	id := c.Param("id")
	parcel, err := services.GetParcelByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No parcel of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	errDel := services.DeleteParcel(id)
	if errDel != nil {
		c.JSON(500, gin.H{
			"error": errDel.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Parcel deleted successfully!",
		"data":    parcel,
	})
}
