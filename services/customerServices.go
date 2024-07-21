package services

import (
	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/repositories"
)

func CreateCustomer(body *models.CustomerReqBody) (*models.Customer, error) {

	// var customer *models.Customer
	customer := &models.Customer{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Contact:   body.Contact,
	}

	// passing memory location of a Customer Struc
	if err := repositories.CreateCustomer(customer); err != nil {
		return nil, err
	}

	return customer, nil
}

func GetAllCustomers() ([]models.Customer, error) {
	return repositories.GetAllCustomers()
}

func GetCustomerByID(id string) (models.Customer, error) {
	return repositories.GetCustomerByID(id)
}

func UpdateCustomerByID(id string, body *models.CustomerReqBody) (models.Customer, error) {
	customer, err := GetCustomerByID(id)
	if err != nil {
		return models.Customer{}, nil
	}

	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Contact = body.Contact

	errUpd := repositories.UpdateCustomerByID(&customer)
	if errUpd != nil {
		return models.Customer{}, errUpd
	}

	return customer, nil
}

func DeleteCustomer(id string) error {
	return repositories.DeleteCustomer(id)
}
