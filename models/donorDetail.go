package models

import (
	"errors"
	"fmt"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/responses"
	"go-donor-list-backend/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DonorDetailCreateRequest struct{
	VolumePerDonation float64        `json:"volumePerDonation" binding:"required"`
	NumberOfDonations int            `json:"numberOfDonations" binding:"required"`
	Type              utils.DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

type DonorDetailUpdateRequest struct{
	ID								string				 `json:"id"`
	VolumePerDonation float64        `json:"volumePerDonation" binding:"required"`
	NumberOfDonations int            `json:"numberOfDonations" binding:"required"`
	Type              utils.DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

type DonorDetail struct {
	ID                string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	VolumePerDonation float64        `json:"volumePerDonation" binding:"required"`
	NumberOfDonations int            `json:"numberOfDonations" binding:"required"`
	Type              utils.DonorType      `json:"type" binding:"required"`
	UserID            string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

// BeforeCreate These functions are called before creating any Post
func (donorDetail *DonorDetail) BeforeCreate(_ *gorm.DB) (err error) {
	donorDetail.ID = uuid.New().String()
	return
}

func (req *DonorDetailCreateRequest) CreateDonorDetail() (responses.DonorDetailResponse, error) {
	fmt.Println("In create-donor-detail, requestPayload : ", req)
	//----> Map donorDetailCreate to DonorDetail
	donorDetail := donorDetailCreateRequestToEntity(req)

	fmt.Println("In create-donor-detail, entity : ", donorDetail)

	//----> Insert the donor-detail into the database.
	if err := initializers.DB.Create(&donorDetail).Error; err != nil {
		return responses.DonorDetailResponse{}, errors.New("failed to create donor detail from database")
	}

	//----> Map donorDetail to donorDetailResponse
	donorDetailResponse := donorDetailEntityToResponse(donorDetail)

	//----> send back the response.
	return donorDetailResponse, nil
}

func (d *DonorDetail) DeleteDonorDetailById(id string, userAuth utils.UserAuth) error {
	//----> Retrieve the donor-detail.
	if _, err := getOneDonorDetail(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the donor-detail from the database.
	if err := initializers.DB.Where("id = ?", id).Delete(&DonorDetail{}).Error; err != nil {
		return errors.New("failed to delete donor detail from database")
	}

	//----> Send back the response.
	return nil
}

func (donorDetail *DonorDetail) DeleteAllDonorDetails() error{
	donorDetails := new([]DonorDetail) //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&donorDetails).Error; err != nil {
		return errors.New("failed to retrieve blood stats")
	}

	//----> Collect all the ids of blood-stats to delete.
	idsOfDonorDetail := getAllDonorDetailIds(*donorDetails)
	
	//----> Delete all blood-stats.
	 if err := initializers.DB.Unscoped().Delete(idsOfDonorDetail).Error; err != nil {
		return errors.New("error deleting blood-stats from database")
	 }

	//----> Send back response.
	return nil
}

func (donorDetail *DonorDetail) DeleteAllDonorDetailsByUserId(userId string) error{
	donorDetails := new([]DonorDetail)
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Find(&donorDetails, DonorDetail{UserID: userId}).Error; err != nil {
		return errors.New(err.Error())
	}
	
	//----> Delete the blood-stat
	if err := initializers.DB.Where(&DonorDetail{UserID: userId}).Delete(&DonorDetail{}).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (req *DonorDetailUpdateRequest) EditDonorDetailById(id string, userAuth utils.UserAuth) error {
	//----> Map donorDetail to donorDetailUpdateRequest
	donorDetail := donorDetailUpdateRequestToEntity(req)

	//----> Retrieve the donor-detail.
	if _, err := getOneDonorDetail(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Update the donor-detail in the database.
	if err := initializers.DB.Model(&donorDetail).Updates(donorDetail).Error; err != nil {
		return errors.New("failed to update donor detail from database")
	}

	//----> Send back the response.
	return nil
}

func (d *DonorDetail) GetDonorDetailByID(id string, userAuth utils.UserAuth) (responses.DonorDetailResponse, error) {
	//----> Retrieve the donor-detail from the database.
	donorDetail, err := getOneDonorDetail(id, userAuth)

	//----> Check for error.
	if err != nil {
		return responses.DonorDetailResponse{}, errors.New(err.Error())
	}

	//----> Map donorDetail to donorDetailResponse.
	donorDetailResponse := donorDetailEntityToResponse(donorDetail)

	//----> send back the response.
	return donorDetailResponse, nil
}

func (d *DonorDetail) GetAllDonorDetails() ([]responses.DonorDetailResponse, error) {
	donors := make([]DonorDetail, 0) //----> Declare the variable.

	//----> Retrieve all donor-details.
	if err := initializers.DB.Find(&donors).Error; err != nil {
		return nil, errors.New("failed to get donor detail from database")
	}

	//----> Map donors to donorsResponse.
	donorsResponse := donorDetailListEntityToListResponse(donors)

	//----> Send back the response.
	return donorsResponse, nil
}

func (d *DonorDetail) GetAllDonorDetailsByUserId(userId string) ([]responses.DonorDetailResponse, error){
	donorDetails := new([]DonorDetail)
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Find(&donorDetails, DonorDetail{UserID: userId}).Error; err != nil {
		return []responses.DonorDetailResponse{},errors.New(err.Error())
	}

	//----> map slice list to vitaResponse slice
	donorDetailsResponse := donorDetailListEntityToListResponse(*donorDetails)


	//---> Send back the response
	return donorDetailsResponse, nil
}

func getOneDonorDetail(id string, userAuth utils.UserAuth) (DonorDetail, error) {
	var donorDetail DonorDetail //----> Declare the variable.

	//----> Retrieve the donor-detail with the given id from database.
	if err := initializers.DB.Where("id = ?", id).First(&donorDetail).Error; err != nil {
		return DonorDetail{}, errors.New("failed to get donor detail from database")
	}

	//----> Check for ownership and admin privilege.
	if err := utils.CheckForOwnership(userAuth.UserId, donorDetail.UserID, userAuth.IsAdmin); err != nil{
		return DonorDetail{}, errors.New("you are not permitted to view or perform any action on this resource")
	}

	//----> Send back the response.
	return donorDetail, nil
}

func getManyDonorDetailByUserId(userId string, donorDetails []DonorDetail)(error){
	//----> Get all donor-details by user-id.
	if err := initializers.DB.Where(&DonorDetail{UserID: userId}).Preload("User").Find(&donorDetails).Error; err != nil{
		return errors.New("donor-details cannot be retrieved from database")
	}

	//----> Send back the response.
	return nil
}

func getAllDonorDetailIds(donorDetails []DonorDetail) []DonorDetail {
	donorDetailIds := []DonorDetail{}

	//----> Collect all the blood-stat ids.
	for _, donorDetail := range donorDetails{
		donorDetail := DonorDetail{ID: donorDetail.ID}

		donorDetailIds = append(donorDetailIds, donorDetail)
	}

	//----> send back the result
	return donorDetailIds
}

func donorDetailCreateRequestToEntity(req *DonorDetailCreateRequest)DonorDetail{
	return DonorDetail{
		VolumePerDonation: req.VolumePerDonation,
		NumberOfDonations: req.NumberOfDonations,
		Type: req.Type,
		UserID: req.UserID,
	}
}
func donorDetailUpdateRequestToEntity(req *DonorDetailUpdateRequest)DonorDetail{
	return DonorDetail{
		ID: req.ID,
		VolumePerDonation: req.VolumePerDonation,
		NumberOfDonations: req.NumberOfDonations,
		Type: req.Type,
		UserID: req.UserID,
	}
}

func donorDetailEntityToResponse(res DonorDetail)responses.DonorDetailResponse{
	return responses.DonorDetailResponse{
		ID: res.ID,
		VolumePerDonation: res.VolumePerDonation,
		NumberOfDonations: res.NumberOfDonations,
		Type: res.Type,
		UserID: res.UserID,
	}
}

func donorDetailListEntityToListResponse(list []DonorDetail)[]responses.DonorDetailResponse {
	listResponse := []responses.DonorDetailResponse{}

	for _, res := range list {
		listResponse = append(listResponse, donorDetailEntityToResponse(res))
	}

	return listResponse
}

