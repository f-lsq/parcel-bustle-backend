package repositories

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func CreateCustomer(customer *models.Customer) error {
	result := config.DB.Create(customer)
	return result.Error
}
