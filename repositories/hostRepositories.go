package repositories

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func CreateHost(host *models.Host) error {
	result := config.DB.Create(host)
	return result.Error
}

func GetAllHosts() ([]models.Host, error) {
	var hosts []models.Host

	result := config.DB.Find(&hosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return hosts, nil
}

func GetHostByID(id string) (models.Host, error) {
	var host models.Host

	result := config.DB.First(&host, id)
	if result.Error != nil {
		return host, result.Error
	}
	return host, nil
}

func UpdateHostByID(host *models.Host) error {
	result := config.DB.Save(host)
	return result.Error
}

func DeleteHost(id string) error {
	result := config.DB.Delete(&models.Host{}, id)
	return result.Error
}
