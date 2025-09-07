package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DonorType string

const (
	FirstTimeDonor DonorType = "FirstTimeDonor"
	FrequentDonor  DonorType = "FrequentDonor"
	OneOfDonor     DonorType = "OneOfDonor"
)

const ()

type DonorDetail struct {
	ID                string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	VolumePerDonation float64        `json:"volume_per_donation" binding:"required"`
	NumberOfDonations int            `json:"numberOfTimes" binding:"required"`
	Type              DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

// BeforeCreate These functions are called before creating any Post
func (donorDetail *DonorDetail) BeforeCreate(_ *gorm.DB) (err error) {
	donorDetail.ID = uuid.New().String()
	return
}

func (donorDetail *DonorDetail) CreateDonorDetail() (DonorDetail, error) {
	//----> Insert the donor-detail into the database.
	if err := initializers.DB.Create(donorDetail).Error; err != nil {
		return DonorDetail{}, errors.New("failed to create donor detail from database")
	}

	//----> send back the response.
	return *donorDetail, nil
}

func (d *DonorDetail) DeleteDonorDetailById(id string) error {
	//----> Retrieve the donor-detail.
	if _, err := getOneDonorDetail(id); err != nil {
		return errors.New("failed to get donor detail from database")
	}

	//----> Delete the donor-detail from the database.
	if err := initializers.DB.Delete(&DonorDetail{}).Error; err != nil {
		return errors.New("failed to delete donor detail from database")
	}

	//----> Send back the response.
	return nil
}

func (donorDetail *DonorDetail) EditDonorDetailById(id string) error {
	//----> Retrieve the donor-detail.
	if _, err := getOneDonorDetail(id); err != nil {
		return errors.New("failed to get donor detail from database")
	}

	//----> Update the donor-detail in the database.
	if err := initializers.DB.Model(&donorDetail).Updates(donorDetail).Error; err != nil {
		return errors.New("failed to update donor detail from database")
	}

	//----> Send back the response.
	return nil
}

func (d *DonorDetail) GetDonorDetailByID(id string) (DonorDetail, error) {
	//----> Retrieve the donor-detail from the database.
	donorDetail, err := getOneDonorDetail(id)

	//----> Check for error.
	if err != nil {
		return DonorDetail{}, errors.New("failed to get donor detail from database")
	}

	//----> send back the response.
	return donorDetail, nil
}

func (d *DonorDetail) GetAllDonorDetails() ([]DonorDetail, error) {
	donors := make([]DonorDetail, 0) //----> Declare the variable.

	//----> Retrieve all donor-details.
	if err := initializers.DB.Find(&donors).Error; err != nil {
		return nil, errors.New("failed to get donor detail from database")
	}

	//----> Send back the response.
	return donors, nil
}

func getOneDonorDetail(id string) (DonorDetail, error) {
	var donorDetail DonorDetail //----> Declare the variable.

	//----> Retrieve the donor-detail with the given id from database.
	if err := initializers.DB.Where("id = ?", id).First(&donorDetail).Error; err != nil {
		return DonorDetail{}, errors.New("failed to get donor detail from database")
	}

	//----> Send back the response.
	return donorDetail, nil
}
