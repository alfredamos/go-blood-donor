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

type BloodStatCreateRequest struct {
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string 				`json:"userId" binding:"required"`
}

type BloodStatUpdateRequest struct {
	ID 				 string         `json:"id"`
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string 				`json:"userId" binding:"required"`
}

type BloodStat struct {
	ID         string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	GenoType   string         `json:"genoType" binding:"required"`
	BloodGroup string         `json:"bloodGroup" binding:"required"`
	UserID     string         
}

// BeforeCreate These functions are called before creating any Post
func (bloodStat *BloodStat) BeforeCreate(_ *gorm.DB) error {
	bloodStat.ID = uuid.New().String()
	return nil
}

func (req *BloodStatCreateRequest) CreateBloodStat() (responses.BloodStatResponse, error) {
	//----> Map BloodStatCreateRequest to BloodStat 
		bloodStat := bloodStatCreateRequestToEntity(req)
	//----> Insert the blood-stat into the database.
	if err := initializers.DB.Create(&bloodStat).Error; err != nil {
		return responses.BloodStatResponse{}, errors.New("failed to create blood stat")
	}

	//----> Map BloodStat to BloodStatResponse
	bloodStatResponse := bloodStateEntityToResponse(bloodStat)
	//----> Send back the response.
	return bloodStatResponse, nil
}

func (b *BloodStat) DeleteBloodStatById(id string, userAuth utils.UserAuth) error {
	//----> retrieve blood-stat and check for error.
	if _, err := getOneBloodStat(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the blood-stat
	if err := initializers.DB.Where("id = ?", id).Delete(&BloodStat{}, "id = ?", id).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (bloodStat *BloodStat) DeleteBloodStatByUserId(userId string) error {
	bloodStats := new([]BloodStat)
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Find(&bloodStats, BloodStat{UserID: userId}).Error; err != nil {
		return errors.New(err.Error())
	}
	
	//----> Delete the blood-stat
	if err := initializers.DB.Where(&BloodStat{UserID: userId}).Delete(&BloodStat{}).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (bloodStat *BloodStat) DeleteAllBloodStat() error {
	bloodStats := new([]BloodStat) //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&bloodStats).Error; err != nil {
		return errors.New("failed to retrieve blood stats")
	}

	//----> Collect all the ids of blood-stats to delete.
	idsOfBloodStat := getAllBloodStatIds(*bloodStats)
	
	//----> Delete all blood-stats.
	 if err := initializers.DB.Unscoped().Delete(idsOfBloodStat).Error; err != nil {
		return errors.New("error deleting blood-stats from database")
	 }

	//----> Send back response.
	return nil
}

func (req *BloodStatUpdateRequest) EditBloodStatById(id string, userAuth utils.UserAuth) error {
	//----> retrieve blood-stat and check for error.
	if _, err := getOneBloodStat(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Map bloodStatUpdate to bloodStat
	bloodStat := bloodStatUpdateRequestToEntity(req)

	//----> Edit the blood-stat
	if err := initializers.DB.Model(&bloodStat).Updates(bloodStat).Error; err != nil {
		return errors.New("failed to update blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (b *BloodStat) GetBloodStatById(id string, userAuth utils.UserAuth) (responses.BloodStatResponse, error) {
	//----> Retrieve the blood-stat from database.
	bloodStat, err := getOneBloodStat(id, userAuth)

	//----> Check for error.
	if err != nil {
		return responses.BloodStatResponse{}, errors.New(err.Error())
	}
	
	//----> Map bloodStat to bloodStatResponse.
	bloodStatResponse := bloodStateEntityToResponse(bloodStat)
	//----> send back the response.
	return bloodStatResponse, nil
}

func (b *BloodStat) GetBloodStatByUserId(userId string)(responses.BloodStatResponse, error){
	var bloodStat BloodStat //----> Declare the variable.
	fmt.Println("Are you in the right-place?")
	//----> Retrieve the blood stat.
	if err := initializers.DB.Where(&BloodStat{UserID: userId}).First(&bloodStat).Error; err != nil {
		return responses.BloodStatResponse{}, errors.New(err.Error())
	}

	//----> Map bloodStat to BloodStatResponse
	bloodStatResponse := bloodStateEntityToResponse(bloodStat)

	//----> Send back response.
	return bloodStatResponse, nil

}


func (b *BloodStat) GetAllBloodStat() ([]responses.BloodStatResponse, error) {
	bloodStats := []BloodStat{} //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&bloodStats).Error; err != nil {
		return []responses.BloodStatResponse{}, errors.New("failed to retrieve blood stats")
	}

	//----> Map slice of bloodStat to bloodStatResponse
	bloodStatsResponse := bloodStatListEntityToListResponse(bloodStats)

	//----> Send back the response.
	return bloodStatsResponse, nil
}

func getOneBloodStat(id string, userAuth utils.UserAuth) (BloodStat, error) {
	var bloodStat BloodStat //----> Declare the variable.

	//----> Retrieve the blood stat.
	if err := initializers.DB.Where("id = ?", id).First(&bloodStat).Error; err != nil {
		return bloodStat, errors.New("failed to retrieve blood stat from database")
	}

	//----> Check for ownership and admin privilege.
	if err := utils.CheckForOwnership(userAuth.UserId, bloodStat.UserID, userAuth.IsAdmin); err != nil{
		return BloodStat{}, errors.New("you are not permitted to view or perform any action on this resource")
	}
	//----> Send back the response
	return bloodStat, nil
}

func getManyBloodStatByUserId(userId string, bloodStats []BloodStat)(error){
	//----> Get all donor-details by user-id.
	if err := initializers.DB.Preload("User").Find(&bloodStats, BloodStat{UserID: userId}).Error; err != nil{
		return errors.New("donor-details cannot be retrieved from database")
	}

	//----> Send back the response.
	return nil
}


func getAllBloodStatIds(bloodStats []BloodStat) []BloodStat {
	fmt.Println("$$$$$$$, in get ids &&&&&, bloodStats : ", bloodStats)
	bloodStatIds := []BloodStat{}

	//----> Collect all the blood-stat ids.
	for _, bloodStat := range bloodStats{
		bloodStat := BloodStat{ID: bloodStat.ID}

		fmt.Println("bloodStat, in loop, : ", bloodStat)

		bloodStatIds = append(bloodStatIds, bloodStat)
	}

	//----> send back the result
	return bloodStatIds
}

func bloodStatCreateRequestToEntity(req *BloodStatCreateRequest)BloodStat{
	return BloodStat{
		GenoType: req.GenoType,
		BloodGroup: req.BloodGroup,
		UserID: req.UserID,
	}
}
func bloodStatUpdateRequestToEntity(req *BloodStatUpdateRequest)BloodStat{
	return BloodStat{
		ID: req.ID,
		GenoType: req.GenoType,
		BloodGroup: req.BloodGroup,
		UserID: req.UserID,
	}
}

func bloodStateEntityToResponse(res BloodStat)responses.BloodStatResponse{
	return responses.BloodStatResponse{
		ID: res.ID,
		BloodGroup: res.BloodGroup,
		GenoType: res.GenoType,
		UserID: res.UserID,
	}
}

func bloodStatListEntityToListResponse(list []BloodStat)[]responses.BloodStatResponse {
	listResponse := []responses.BloodStatResponse{}

	for _, res := range list {
		listResponse = append(listResponse, bloodStateEntityToResponse(res))
	}

	return listResponse
}
