package main

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Worker{}) // migrate 'Worker' from 'models' package
	config.DB.AutoMigrate(&models.Parcel{}) // migrate 'Parcel' from 'models' package
}
