package controllers

import (
	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/services"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var body models.CustomerReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer, err := services.CreateCustomer(&body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Customer created successfully",
		"data":    customer,
	})
}

func GetCustomers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer retrieved",
	})
}

func GetCustomerByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer retrieved by ID",
	})
}

func UpdateCustomer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer updated",
	})
}

func DeleteCustomer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Customer deleted",
	})
}
