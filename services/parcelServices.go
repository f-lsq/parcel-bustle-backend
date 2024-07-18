package services

import (
	"time"

	"github.com/f-lsq/parcel-bustle-backend/models"
	"github.com/f-lsq/parcel-bustle-backend/repositories"
)

func CreateParcel(body models.ParcelReqBody) (*models.Parcel, error) {
	var deliverByDate time.Time
	var err error

	if body.DeliverBy == "" {
		deliverByDate = time.Now().AddDate(0, 0, 7)
	} else {
		deliverByDate, err = time.Parse("2006-01-02", body.DeliverBy)
		if err != nil {
			return nil, err
		}
	}

	// var parcel *models.Parcel
	parcel := &models.Parcel{
		Status:          body.Status,
		DeliveryAddress: body.DeliveryAddress,
		ReturnAddress:   body.ReturnAddress,
		DeliverBy:       deliverByDate,
		DeliveredImage:  body.DeliveredImage,
		PaymentType:     body.PaymentType,
		PaymentMade:     body.PaymentMade,
	}

	// passing memory location of a Parcel Struc
	if err := repositories.CreateParcel(parcel); err != nil {
		return nil, err
	}

	return parcel, nil
}

func GetAllParcels() ([]models.Parcel, error) {
	return repositories.GetAllParcels()
}

func GetParcelByID(id string) (models.Parcel, error) {
	return repositories.GetParcelByID(id)
}

func UpdateParcelByID(id string, body models.ParcelReqBody) (models.Parcel, error) {
	parcel, err := GetParcelByID(id)
	if err != nil {
		return models.Parcel{}, err
	}

	var DeliverByDate time.Time
	if body.DeliverBy == "" {
		DeliverByDate = time.Now().AddDate(0, 0, 7)
	} else {
		DeliverByDate, _ = time.Parse("2006-01-02", body.DeliverBy)
	}

	parcel.Status = body.Status
	parcel.DeliveryAddress = body.DeliveryAddress
	parcel.ReturnAddress = body.ReturnAddress
	parcel.DeliverBy = DeliverByDate
	parcel.DeliveredImage = body.DeliveredImage
	parcel.PaymentType = body.PaymentType
	parcel.PaymentMade = body.PaymentMade

	errUpd := repositories.UpdateParcelByID(&parcel)
	if errUpd != nil {
		return models.Parcel{}, errUpd
	}

	return parcel, nil
}

func DeleteParcel(id string) error {
	return repositories.DeleteParcel(id)
}
