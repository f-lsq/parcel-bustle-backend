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
