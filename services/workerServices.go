package services

import (
	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/repositories"
)

func CreateWorker(body *models.WorkerReqBody) (*models.Worker, error) {

	// var worker *models.Worker
	worker := &models.Worker{
		Username:     body.Username,
		Email:        body.Email,
		Password:     body.Password,
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Contact:      body.Contact,
		ProfileImage: body.ProfileImage,
		Verified:     body.Verified,
		Approved:     body.Approved,
	}

	// passing memory location of a Worker Struc
	if err := repositories.CreateWorker(worker); err != nil {
		return nil, err
	}

	return worker, nil
}

func GetAllWorkers() ([]models.Worker, error) {
	return repositories.GetAllWorkers()
}

func GetWorkerByID(id string) (models.Worker, error) {
	return repositories.GetWorkerByID(id)
}

func UpdateWorkerByID(id string, body *models.WorkerReqBody) (models.Worker, error) {
	worker, err := GetWorkerByID(id)
	if err != nil {
		return models.Worker{}, err
	}

	worker.Username = body.Username
	worker.Email = body.Email
	worker.Password = body.Password
	worker.FirstName = body.FirstName
	worker.LastName = body.LastName
	worker.Contact = body.Contact
	worker.ProfileImage = body.ProfileImage
	worker.Verified = body.Verified
	worker.Approved = body.Approved

	errUpd := repositories.UpdateWorkerByID(&worker)
	if errUpd != nil {
		return models.Worker{}, errUpd
	}

	return worker, nil
}

func DeleteWorker(id string) error {
	return repositories.DeleteWorker(id)
}
