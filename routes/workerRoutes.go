package routes

import (
	"github.com/f-lsq/parcel-bustle-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Worker(route *gin.Engine) {
	workers := route.Group("/api/workers")
	{
		workers.POST("/", controllers.CreateWorker)
		workers.GET("/", controllers.GetWorkers)
		workers.GET("/:id", controllers.GetWorkerByID)
		workers.PUT("/:id", controllers.UpdateWorker)
		workers.DELETE("/:id", controllers.DeleteWorker)
	}
}
