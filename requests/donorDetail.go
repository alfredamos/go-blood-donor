package requests

import (
	"go-donor-list-backend/utils"
)

type DonorDetailCreateRequest struct{
	VolumePerDonation float64        `json:"volume_per_donation" binding:"required"`
	NumberOfDonations int            `json:"numberOfTimes" binding:"required"`
	Type              utils.DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

type DonorDetailUpdateRequest struct{
	ID								string				 `json:"id"`
	VolumePerDonation float64        `json:"volume_per_donation" binding:"required"`
	NumberOfDonations int            `json:"numberOfTimes" binding:"required"`
	Type              utils.DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}