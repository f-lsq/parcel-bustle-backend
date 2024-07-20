package routes

import (
	"github.com/f-lsq/parcel-bustle-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Customer(route *gin.Engine) {
	customers := route.Group("/api/customers")
	{
		customers.POST("/", controllers.CreateCustomer)
		customers.GET("/", controllers.GetCustomers)
		customers.GET("/:id", controllers.GetCustomerByID)
		customers.PUT("/:id", controllers.UpdateCustomer)
		customers.DELETE("/:id", controllers.DeleteCustomer)
	}
}
