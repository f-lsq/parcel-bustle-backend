package repositories

import (
	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func CreateWorker(worker *models.Worker) error {
	result := config.DB.Create(worker)
	return result.Error
}

func GetAllWorkers() ([]models.Worker, error) {
	var workers []models.Worker

	result := config.DB.Find(&workers)
	if result.Error != nil {
		return nil, result.Error
	}
	return workers, nil
}

func GetWorkerByID(id string) (models.Worker, error) {
	var worker models.Worker

	result := config.DB.First(&worker, id)
	if result.Error != nil {
		return worker, result.Error
	}
	return worker, nil
}

func UpdateWorkerByID(worker *models.Worker) error {
	result := config.DB.Save(worker)
	return result.Error
}

func DeleteWorker(id string) error {
	result := config.DB.Delete(&models.Worker{}, id)
	return result.Error
}
