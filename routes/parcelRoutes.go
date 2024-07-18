package routes

import (
	"github.com/f-lsq/parcel-bustle-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Parcel(route *gin.Engine) {
	parcels := route.Group("/api/parcels")
	{
		parcels.POST("/", controllers.CreateParcel)
		parcels.GET("/", controllers.GetParcels)
		parcels.GET("/:id", controllers.GetParcelByID)
		parcels.PUT("/:id", controllers.UpdateParcel)
		parcels.DELETE("/:id", controllers.DeleteParcel)
	}
}
