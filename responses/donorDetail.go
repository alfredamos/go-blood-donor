package responses

import "go-donor-list-backend/utils"

type DonorDetailResponse struct {
	ID                string          `json:"id"`
	VolumePerDonation float64         `json:"volumePerDonation" binding:"required"`
	NumberOfDonations int             `json:"numberOfDonations" binding:"required"`
	Type              utils.DonorType `json:"type" binding:"required"`
	UserID            string          `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}