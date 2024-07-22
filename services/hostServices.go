package services

import (
	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/repositories"
)

func CreateHost(body *models.HostReqBody) (*models.Host, error) {

	// var host *models.Host
	host := &models.Host{
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	// passing memory location of a Host Struc
	if err := repositories.CreateHost(host); err != nil {
		return nil, err
	}

	return host, nil
}

func GetAllHosts() ([]models.Host, error) {
	return repositories.GetAllHosts()
}

func GetHostByID(id string) (models.Host, error) {
	return repositories.GetHostByID(id)
}

func UpdateHostByID(id string, body *models.HostReqBody) (models.Host, error) {
	host, err := GetHostByID(id)
	if err != nil {
		return models.Host{}, err
	}

	host.Username = body.Username
	host.Email = body.Email
	host.Password = body.Password
	host.FirstName = body.FirstName
	host.LastName = body.LastName

	errUpd := repositories.UpdateHostByID(&host)
	if errUpd != nil {
		return models.Host{}, errUpd
	}

	return host, nil
}

func DeleteHost(id string) error {
	return repositories.DeleteHost(id)
}
