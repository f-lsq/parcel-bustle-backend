package routes

import (
	"github.com/f-lsq/parcel-bustle-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Host(route *gin.Engine) {
	hosts := route.Group("/api/hosts")
	{
		hosts.POST("/", controllers.CreateHost)
		hosts.GET("/", controllers.GetHosts)
		hosts.GET("/:id", controllers.GetHostByID)
		hosts.PUT("/:id", controllers.UpdateHost)
		hosts.DELETE("/:id", controllers.DeleteHost)
	}
}
