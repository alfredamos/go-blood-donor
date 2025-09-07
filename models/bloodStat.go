package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (bloodStat *BloodStat) CreateBloodStat() (BloodStat, error) {
	//----> Insert the blood-stat into the database.
	if err := initializers.DB.Create(&bloodStat).Error; err != nil {
		return BloodStat{}, errors.New("failed to create blood stat")
	}

	//----> Send back the response.
	return *bloodStat, nil
}

func (_ *BloodStat) DeleteBloodStatById(id string) error {
	//----> retrieve blood-stat and check for error.
	if _, err := getOneBloodStat(id); err != nil {
		return errors.New("failed to retrieve blood stat from database")
	}

	//----> Delete the blood-stat
	if err := initializers.DB.Delete(&BloodStat{}, "id = ?", id).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (bloodStat *BloodStat) EditBloodStatById(id string) error {
	//----> retrieve blood-stat and check for error.
	if _, err := getOneBloodStat(id); err != nil {
		return errors.New("failed to retrieve blood stat from database")
	}

	//----> Edit the blood-stat
	if err := initializers.DB.Model(&bloodStat).Updates(&bloodStat).Error; err != nil {
		return errors.New("failed to update blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (_ *BloodStat) GetBloodStatById(id string) (BloodStat, error) {
	//----> Retrieve the blood-stat from database.
	bloodStat, err := getOneBloodStat(id)

	//----> Check for error.
	if err != nil {
		return BloodStat{}, errors.New("failed to retrieve blood stat from database")
	}

	//----> send back the response.
	return bloodStat, nil
}

func (_ *BloodStat) GetAllBloodStat() ([]BloodStat, error) {
	var bloodStats []BloodStat //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&bloodStats).Error; err != nil {
		return []BloodStat{}, errors.New("failed to retrieve blood stats")
	}

	//----> Send back the response.
	return bloodStats, nil
}

func getOneBloodStat(id string) (BloodStat, error) {
	var bloodStat BloodStat //----> Declare the variable.

	//----> Retrieve the blood stat.
	if err := initializers.DB.Where("id = ?", id).First(&bloodStat).Error; err != nil {
		return bloodStat, errors.New("failed to retrieve blood stat from database")
	}

	//----> Send back the response
	return bloodStat, nil
}
