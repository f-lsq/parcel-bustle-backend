package main

import (
	"os"
	"time"

	"github.com/f-lsq/parcel-bustle-backend/config"
	"github.com/f-lsq/parcel-bustle-backend/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	// Create Tables for all Entities
	config.DB.AutoMigrate(&models.Worker{})           // migrate 'Worker' from 'models' package
	config.DB.AutoMigrate(&models.Parcel{})           // migrate 'Parcel' from 'models' package
	config.DB.AutoMigrate(&models.Customer{})         // migrate 'Customer' from 'models' package
	config.DB.AutoMigrate(&models.Host{})             // migrate 'Host' from 'models' package
	config.DB.AutoMigrate(&models.BlacklistedToken{}) // migrate 'BlacklistedToken' from 'models' package

	// Seed data for all Entities
	seedWorkers := []models.Worker{
		{Username: "mulanhua", Email: "mulanhua@parcelbustle.com", Password: "mulanhua123@", FirstName: "Mulan", LastName: "Hua", Contact: "+65 9686 2975", ProfileImage: "", Verified: true, Approved: true},
		{Username: "shangli", Email: "shangli@parcelbustle.com", Password: "shangli123@", FirstName: "Shang", LastName: "Li", Contact: "+65 9738 1068", ProfileImage: "", Verified: true, Approved: true},
		{Username: "chifu", Email: "chifu@parcelbustle.com", Password: "chifu123@", FirstName: "Chi", LastName: "Fu", Contact: "+65 9898 2784", ProfileImage: "", Verified: true, Approved: false},
	}

	seedParcels := []models.Parcel{
		{WorkerID: 1, CustomerID: 1, Status: "Uncollected", DeliveryAddress: "Aljunied Industrial Complex 623 Aljunied Road #05-03", ReturnAddress: "#03-01 Cogent Logistics Hub 1 Buroh Crescent #6M-01, 627545", DeliverBy: time.Now().AddDate(0, 0, 7), PaymentType: "Cash", PaymentMade: false},
		{WorkerID: 1, CustomerID: 2, Status: "Collected", DeliveryAddress: "1 Boon Leat Terrace, #07-02, Harbourside Building 1", ReturnAddress: "#03-01 Cogent Logistics Hub 1 Buroh Crescent #6M-01, 627545", DeliverBy: time.Now().AddDate(0, 0, 7), PaymentType: "Card", PaymentMade: true},
		{WorkerID: 2, CustomerID: 2, Status: "Out for delivery", DeliveryAddress: "Lippo Centre 78 Shenton Way #13-03", ReturnAddress: "#03-01 Cogent Logistics Hub 1 Buroh Crescent #6M-01, 627545", DeliverBy: time.Now().AddDate(0, 0, 7), PaymentType: "PayNow", PaymentMade: false},
		{WorkerID: 2, CustomerID: 3, Status: "Delivered", DeliveryAddress: "HDB Yishun 106 Yishun Ring Road #01-197", ReturnAddress: "#03-01 Cogent Logistics Hub 1 Buroh Crescent #6M-01, 627545", DeliverBy: time.Now().AddDate(0, 0, 7), PaymentType: "Cash", PaymentMade: true},
		{WorkerID: 2, CustomerID: 4, Status: "Undelivered", DeliveryAddress: "3017 Bedok Nth St 5 #05-02", ReturnAddress: "#03-01 Cogent Logistics Hub 1 Buroh Crescent #6M-01, 627545", DeliverBy: time.Now().AddDate(0, 0, 7), PaymentType: "Card", PaymentMade: false},
	}

	seedCustomers := []models.Customer{
		{FirstName: "Eric", LastName: "Tan", Contact: "+65 9583 9257"},
		{FirstName: "Jia Hui", LastName: "Lim", Contact: "+65 9766 5123"},
		{FirstName: "Jun Jie", LastName: "Li", Contact: "+65 9749 7488"},
		{FirstName: "Samantha", LastName: "", Contact: "+65 9448 8542"},
	}

	seedHosts := []models.Host{
		{Username: "host1", Email: "host1@parcelbustle.com", Password: os.Getenv("HOST1_PASS"), FirstName: "Host 1", LastName: ""},
		{Username: "host2", Email: "host2@parcelbustle.com", Password: os.Getenv("HOST2_PASS"), FirstName: "Host 2", LastName: ""},
		{Username: "host3", Email: "host3@parcelbustle.com", Password: os.Getenv("HOST3_PASS"), FirstName: "Host 3", LastName: ""},
	}

	// Insert seed data for all Entities
	for _, worker := range seedWorkers {
		config.DB.Create(&worker)
	}

	for _, customer := range seedCustomers {
		config.DB.Create(&customer)
	}

	for _, parcel := range seedParcels {
		config.DB.Create(&parcel)
	}

	for _, host := range seedHosts {
		config.DB.Create(&host)
	}
}
