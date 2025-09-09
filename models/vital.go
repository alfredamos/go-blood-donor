package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (vital *Vital) CreateVital() (Vital, error) {
	//----> Calculate the body mass index.
	bodyMassIndex := calculateBMI(vital.Weight, vital.Height)
	vital.BMI = bodyMassIndex

	//----> Insert the new vital into the database.
	if err := initializers.DB.Create(&vital).Error; err != nil {
		return Vital{}, errors.New("failed to create Vital")
	}

	//----> Send back the response.
	return *vital, nil
}

func (v *Vital) DeleteVitalById(id string, userAuth utils.UserAuth) error {
	//----> Retrieve the vital with the given id.
	if _, err := getOneVital(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the vital with the given id
	if err := initializers.DB.Where("id = ?", id).Delete(&Vital{}, "id = ?", id).Error; err != nil {
		return errors.New("failed to delete Vital")
	}

	//----> Send back the response.
	return nil
}

func (v *Vital) DeleteAllVitals()error{
	vitals := new([]Vital)

	//----> Retrieve the vitals from database.
	if err := initializers.DB.Find(&vitals).Error; err != nil {
		return errors.New("failed to retrieve Vital from database")
	}

	//----> Get all the ids of vitals to delete.
	idsOfVitalsToDelete := getIdsOfVitalsToDelete(*vitals)

	//----> Delete all the vitals.
	if err := initializers.DB.Delete(&idsOfVitalsToDelete); err != nil {
		return errors.New("vitals cannot be deleted from database")
	}

	//----> Send back the result.
	return nil
}
func (v *Vital) DeleteAllVitalsByUserId(userId string)error{
	vitals := new([]Vital)

	//----> Retrieve the vitals from database.
	if err := getManyVitalsByUserId(userId, *vitals); err != nil {
		return errors.New("failed to retrieve Vital from database")
	}

	//----> Get all the ids of vitals to delete.
	idsOfVitalsToDelete := getIdsOfVitalsToDelete(*vitals)

	//----> Delete all the vitals.
	if err := initializers.DB.Delete(&idsOfVitalsToDelete); err != nil {
		return errors.New("vitals cannot be deleted from database")
	}

	//----> Send back the result.
	return nil
}

func (vital *Vital) EditVitalById(id string, userAuth utils.UserAuth) error {
	//----> Retrieve the vital with the given id.
	if _, err := getOneVital(id, userAuth); err != nil {
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

func (d *Vital) GetVitalById(id string, userAuth utils.UserAuth) (Vital, error) {
	//----> Retrieve the vital from the database.
	vital, err := getOneVital(id, userAuth)

	//----> Check for error.
	if err != nil {
		return Vital{}, errors.New(err.Error())
	}

	//----> send back the response.
	return vital, nil

}

func (d *Vital) GetAllVitals() ([]Vital, error) {
	var vitals []Vital //----> Declare a slice of vitals.

	//----> Retrieve the vitals from database.
	if err := initializers.DB.Find(&vitals).Error; err != nil {
		return vitals, errors.New("failed to retrieve Vital from database")
	}

	//----> Send back the response.
	return vitals, nil
}

func (v *Vital) GetAllVitalsByUserId(userId string)([]Vital, error){
	vitals := new([]Vital)

	//----> Retrieve all vitals by user-id.
	if err := getManyVitalsByUserId(userId, *vitals); err != nil {
		return []Vital{}, errors.New("vitals for this user cannot be retrieved")
	}

	//---> Send back the response
	return *vitals, nil
}

func getOneVital(id string, userAuth utils.UserAuth) (Vital, error) {
	var vital Vital //----> Declare the variable.

	//----> Retrieve the vital with given id.
	if err := initializers.DB.First(&vital, "id = ?", id).Error; err != nil {
		return vital, errors.New("failed to retrieve Vital from DB")
	}

	//----> Check for ownership and admin privilege.
	if err := utils.CheckForOwnership(userAuth.UserId, vital.UserID, userAuth.IsAdmin); err != nil{
		return Vital{}, errors.New("you are not permitted to view or perform any action on this resource")
	}

	//----> Send back the response.
	return vital, nil
}

func getManyVitalsByUserId(userId string, vitals []Vital) error{
	//----> Retrieve all vitals by user-id.
	if err := initializers.DB.Preload("User").Find(&vitals, Vital{UserID: userId}); err != nil {
		return errors.New("vitals for this user cannot be retrieved")
	}

	//----> Send back the result.
	return nil
}

func getIdsOfVitalsToDelete(vitals []Vital)[]Vital {
	idsOfVitalsToDelete := []Vital{} 

	//----> Collect all the ids of vital to delete
	for _, vital := range vitals {
		idOfVital := Vital{ID: vital.ID}
		idsOfVitalsToDelete = append(idsOfVitalsToDelete, idOfVital)
	}

	//----> Send back the result.
	return idsOfVitalsToDelete
}

func calculateBMI(weight, height float64) float64{
	return (weight/(height * height))
}
