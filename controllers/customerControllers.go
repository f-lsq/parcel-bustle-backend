package controllers

import (
	"errors"
	"fmt"

	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	customers, err := services.GetAllCustomers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "All customers retrieved successfully",
		"data":    customers,
	})
}

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	customer, err := services.GetCustomerByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No customer of ID %s was found", id),
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
		"message": "Customer retrieved successfully",
		"data":    customer,
	})
}

func UpdateCustomer(c *gin.Context) {
	var body models.CustomerReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	customer, err := services.UpdateCustomerByID(id, &body)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No customer of ID %s was found", id),
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
		"message": "Customer updated successfully",
		"data":    customer,
	})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := services.GetCustomerByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No customer of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errDel := services.DeleteCustomer(id); errDel != nil {
		c.JSON(500, gin.H{
			"error": errDel.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Customer deleted",
		"data":    customer,
	})
}
