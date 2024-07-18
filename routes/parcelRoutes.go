package routes

import (
	"github.com/f-lsq/parcel-bustle-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Parcel(route *gin.Engine) {
	parcels := route.Group("/api/parcels")
	{
		parcels.POST("/", controllers.CreateParcel)
	}
}
