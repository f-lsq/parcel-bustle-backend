package controllers

import (
	"errors"
	"fmt"

	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateHost(c *gin.Context) {
	var body models.HostReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	host, err := services.CreateHost(&body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Host created successfully",
		"data":    host,
	})
}

func GetHosts(c *gin.Context) {
	hosts, err := services.GetAllHosts()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "All hosts retrieved successfully",
		"data":    hosts,
	})
}

func GetHostByID(c *gin.Context) {
	id := c.Param("id")
	host, err := services.GetHostByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No host of ID %s was found", id),
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
		"message": "Host retrieved successfully",
		"data":    host,
	})
}

func UpdateHost(c *gin.Context) {
	var body models.HostReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	host, err := services.UpdateHostByID(id, &body)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No host of ID %s was found", id),
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
		"message": "Host updated successfully",
		"data":    host,
	})
}

func DeleteHost(c *gin.Context) {
	id := c.Param("id")

	host, err := services.GetHostByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No host of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errDel := services.DeleteHost(id); errDel != nil {
		c.JSON(500, gin.H{
			"error": errDel.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Host deleted",
		"data":    host,
	})
}
