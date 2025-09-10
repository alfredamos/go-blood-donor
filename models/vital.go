package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/responses"
	"go-donor-list-backend/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VitalCreateRequest struct{
	PressureUp  float64        `json:"pressureUp" binding:"required"`
	PressureLow float64        `json:"pressureLow" binding:"required"`
	Temperature float64        `json:"temperature" binding:"required"`
	Height      float64        `json:"height" binding:"required"`
	Weight      float64        `json:"weight" binding:"required"`
	BMI         float64        `json:"bmi"`
	UserID      string         `json:"userId"`
}

type VitalUpdateRequest struct{
	ID					string         `json:"id"`
	PressureUp  float64        `json:"pressureUp" binding:"required"`
	PressureLow float64        `json:"pressureLow" binding:"required"`
	Temperature float64        `json:"temperature" binding:"required"`
	Height      float64        `json:"height" binding:"required"`
	Weight      float64        `json:"weight" binding:"required"`
	BMI         float64        `json:"bmi"`
	UserID      string         `json:"userId"`
}

type Vital struct {
	ID          string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	PressureUp  float64        `json:"pressureUp" binding:"required"`
	PressureLow float64        `json:"pressureLow" binding:"required"`
	Temperature float64        `json:"temperature" binding:"required"`
	Height      float64        `json:"height" binding:"required"`
	Weight      float64        `json:"weight" binding:"required"`
	BMI         float64        `json:"bmi"`
	UserID      string         `gorm:"foreignKey:UserID;type:varchar(255)" json:"userId" binding:"required"`
}

// BeforeCreate These functions are called before creating any Post
func (vital *Vital) BeforeCreate(_ *gorm.DB) (err error) {
	vital.ID = uuid.New().String()
	return
}

func (req *VitalCreateRequest) CreateVital() (responses.VitalResponse, error) {
	//----> Map vitalCreateRequest to vital
	vital := vitalCreateRequestToEntity(req)
	//----> Calculate the body mass index.
	bodyMassIndex := calculateBMI(vital.Weight, vital.Height)
	vital.BMI = bodyMassIndex

	//----> Insert the new vital into the database.
	if err := initializers.DB.Create(&vital).Error; err != nil {
		return responses.VitalResponse{}, errors.New("failed to create Vital")
	}

	//----> Map vital to vitalResponse.
	vitalResponse := vitalEntityToResponse(vital)

	//----> Send back the response.
	return vitalResponse, nil
}

func (v *Vital) DeleteVitalById(id string, vitalAuth utils.UserAuth) error {
	//----> Retrieve the vital with the given id.
	if _, err := getOneVital(id, vitalAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the vital with the given id
	if err := initializers.DB.Where("id = ?", id).Delete(&Vital{}, "id = ?", id).Error; err != nil {
		return errors.New("failed to delete Vital")
	}

	//----> Send back the response.
	return nil
}

func (v *Vital) DeleteAllVitals() error{
	vitals := new([]Vital) //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&vitals).Error; err != nil {
		return errors.New("failed to retrieve blood stats")
	}

	//----> Collect all the ids of blood-stats to delete.
	idsOfVital := getAllVitalIds(*vitals)
	
	//----> Delete all blood-stats.
	 if err := initializers.DB.Unscoped().Delete(idsOfVital).Error; err != nil {
		return errors.New("error deleting blood-stats from database")
	 }

	//----> Send back response.
	return nil
}

func (vital *Vital) DeleteAllVitalsByUserId(userId string) error{
	vitals := new([]Vital)
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Find(&vitals, Vital{UserID: userId}).Error; err != nil {
		return errors.New(err.Error())
	}
	
	//----> Delete the blood-stat
	if err := initializers.DB.Where(&Vital{UserID: userId}).Delete(&Vital{}).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (req *VitalUpdateRequest) EditVitalById(id string, vitalAuth utils.UserAuth) error {
	//----> map vitalUpdateRequest to vital
	vital := vitalUpdateRequestToEntity(req)

	//----> Retrieve the vital with the given id.
	if _, err := getOneVital(id, vitalAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Calculate the body mass index.
	bodyMassIndex := calculateBMI(vital.Weight, vital.Height)
	vital.BMI = bodyMassIndex

	//----> Update the vital with the given id
	if err := initializers.DB.Model(&vital).Updates(vital).Error; err != nil {
		return errors.New("failed to update Vital")
	}

	//----> Send back the response.
	return nil
}

func (d *Vital) GetVitalById(id string, vitalAuth utils.UserAuth) (responses.VitalResponse, error) {
	//----> Retrieve the vital from the database.
	vital, err := getOneVital(id, vitalAuth)

	//----> Check for error.
	if err != nil {
		return responses.VitalResponse{}, errors.New(err.Error())
	}

	//----> Map vital to vitalResponse
	vitalResponse := vitalEntityToResponse(vital)

	//----> send back the response.
	return vitalResponse, nil

}

func (d *Vital) GetAllVitals() ([]responses.VitalResponse, error) {
	var vitals []Vital //----> Declare a slice of vitals.

	//----> Retrieve the vitals from database.
	if err := initializers.DB.Find(&vitals).Error; err != nil {
		return []responses.VitalResponse{}, errors.New("failed to retrieve Vital from database")
	}

	//----> map slice list to vitaResponse slice
	vitalsResponse := vitalListEntityToListResponse(vitals)

	//----> Send back the response.
	return vitalsResponse, nil
}

func (v *Vital) GetAllVitalsByUserId(userId string)([]responses.VitalResponse, error){
	vitals := new([]Vital)
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Find(&vitals, Vital{UserID: userId}).Error; err != nil {
		return []responses.VitalResponse{},errors.New(err.Error())
	}

	//----> map slice list to vitaResponse slice
	vitalsResponse := vitalListEntityToListResponse(*vitals)


	//---> Send back the response
	return vitalsResponse, nil
}

func getOneVital(id string, userAuth utils.UserAuth) (Vital, error) {
	var vital Vital //----> Declare the variable.

	//----> Retrieve the vital with given id.
	if err := initializers.DB.First(&vital, "id = ?", id).Error; err != nil {
		return vital, errors.New("failed to retrieve Vital from DB")
	}

	//----> Check for ownership and admin privilege.
	if err := utils.CheckForOwnership(userAuth.UserId, vital.ID, userAuth.IsAdmin); err != nil{
		return Vital{}, errors.New("you are not permitted to view or perform any action on this resource")
	}

	//----> Send back the response.
	return vital, nil
}

func getManyVitalsByUserId(userId string, vitals []Vital)(error){
	//----> Get all donor-details by user-id.
	if err := initializers.DB.Where(&Vital{UserID: userId}).Preload("User").Find(&vitals).Error; err != nil{
		return errors.New("donor-details cannot be retrieved from database")
	}

	//----> Send back the response.
	return nil
}

func getAllVitalIds(vitals []Vital) []Vital {
	vitalIds := []Vital{}

	//----> Collect all the blood-stat ids.
	for _, vital := range vitals{
		vital := Vital{ID: vital.ID}

		vitalIds = append(vitalIds, vital)
	}

	//----> send back the result
	return vitalIds
}

func calculateBMI(weight, height float64) float64{
	return (weight/(height * height))
}


func vitalCreateRequestToEntity(req *VitalCreateRequest)Vital{
	return Vital{
		PressureUp  : req.PressureUp,
		PressureLow : req.PressureLow,
		Temperature : req.Temperature,
		Height 			: req.Height,
		Weight      : req.Weight,
		BMI         : req.BMI,
		UserID      : req.UserID, 
	}
}
func vitalUpdateRequestToEntity(req *VitalUpdateRequest)Vital{
	return Vital{
		ID: req.ID,
		PressureUp  : req.PressureUp,
		PressureLow : req.PressureLow,
		Temperature : req.Temperature,
		Height 			: req.Height,
		Weight      : req.Weight,
		BMI         : req.BMI,
		UserID      : req.UserID, 
	}
}

func vitalEntityToResponse(res Vital)responses.VitalResponse{
	return responses.VitalResponse{
		ID: res.ID,
		PressureUp  : res.PressureUp,
		PressureLow : res.PressureLow,
		Temperature : res.Temperature,
		Height 			: res.Height,
		Weight      : res.Weight,
		BMI         : res.BMI,
		UserID      : res.UserID, 
	}
}

func vitalListEntityToListResponse(list []Vital)[]responses.VitalResponse {
	listResponse := []responses.VitalResponse{}

	for _, res := range list {
		listResponse = append(listResponse, vitalEntityToResponse(res))
	}

	return listResponse
}
