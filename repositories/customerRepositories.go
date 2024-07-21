package repositories

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func CreateCustomer(customer *models.Customer) error {
	result := config.DB.Create(customer)
	return result.Error
}

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer

	result := config.DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func GetCustomerByID(id string) (models.Customer, error) {
	var customer models.Customer

	result := config.DB.First(&customer, id)
	if result.Error != nil {
		return customer, result.Error
	}
	return customer, nil
}

func UpdateCustomerByID(customer *models.Customer) error {
	result := config.DB.Save(customer)
	return result.Error
}

func DeleteCustomer(id string) error {
	result := config.DB.Delete(&models.Customer{}, id)
	return result.Error
}
