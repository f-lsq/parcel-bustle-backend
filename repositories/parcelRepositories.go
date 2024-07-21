package repositories

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func CreateParcel(parcel *models.Parcel) error {
	result := config.DB.Create(parcel)
	return result.Error
}

func GetAllParcels() ([]models.Parcel, error) {
	var parcels []models.Parcel

	result := config.DB.Find(&parcels)
	if result.Error != nil {
		return nil, result.Error
	}
	return parcels, nil
}

func GetParcelByID(id string) (models.Parcel, error) {
	var parcel models.Parcel

	result := config.DB.First(&parcel, id)
	if result.Error != nil {
		return parcel, result.Error
	}
	return parcel, nil
}

func UpdateParcelByID(parcel *models.Parcel) error {
	result := config.DB.Save(parcel)
	return result.Error
}

func DeleteParcel(id string) error {
	result := config.DB.Delete(&models.Parcel{}, id)
	return result.Error
}
