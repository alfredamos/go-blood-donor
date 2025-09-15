package main

import (
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/models"
)

func init() {
	//----> Get all environment variables.
	initializers.LoadEnvVariable()

	//----> Connect to database.
	initializers.ConnectDB()
}

func main() {
	//----> Migrate the gorm models into mysql database.
	err := initializers.DB.AutoMigrate(&models.User{}, &models.BloodStat{}, &models.DonorDetail{}, &models.Vital{}, &models.Token{})
	if err != nil {
		return
	}

}
