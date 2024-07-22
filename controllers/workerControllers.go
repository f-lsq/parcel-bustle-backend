package controllers

import (
	"errors"
	"fmt"

	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorker(c *gin.Context) {
	var body models.WorkerReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	worker, err := services.CreateWorker(&body)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "Worker created successfully",
		"data":    worker,
	})
}

func GetWorkers(c *gin.Context) {
	workers, err := services.GetAllWorkers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "All workers retrieved successfully",
		"data":    workers,
	})
}

func GetWorkerByID(c *gin.Context) {
	id := c.Param("id")
	worker, err := services.GetWorkerByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No worker of ID %s was found", id),
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
		"message": "Worker retrieved successfully",
		"data":    worker,
	})
}

func UpdateWorker(c *gin.Context) {
	var body models.WorkerReqBody

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	worker, err := services.UpdateWorkerByID(id, &body)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No worker of ID %s was found", id),
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
		"message": "Worker updated successfully",
		"data":    worker,
	})
}

func DeleteWorker(c *gin.Context) {
	id := c.Param("id")

	worker, err := services.GetWorkerByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{
			"error": fmt.Sprintf("No worker of ID %s was found", id),
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errDel := services.DeleteWorker(id); errDel != nil {
		c.JSON(500, gin.H{
			"error": errDel.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Worker deleted",
		"data":    worker,
	})
}
