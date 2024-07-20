package main

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	router := gin.Default() //initialises a new Gin router

	routes.Parcel(router)
	routes.Customer(router)

	router.Run()
}
