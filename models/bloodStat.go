package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/utils"
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
	//----> retrieve blood-stat and check for error.
	if err := initializers.DB.Where("userId = ?", userId).First(&bloodStat).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the blood-stat
	if err := initializers.DB.Where("userId = ?", userId).Delete(&BloodStat{}, "userId = ?", userId).Error; err != nil {
		return errors.New("failed to delete blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (bloodStat *BloodStat) DeleteAllBloodStat() error {
	bloodStats := new([]BloodStat)
	//----> Retrieve all the bloodStats
	if err := initializers.DB.Find(&bloodStat); err != nil {
		return errors.New("blood-stats cannot be retrieved from database")
	}

	//----> Get all the blood-stat ids
	bloodStatIds := getAllBloodStatIds(*bloodStats)

	//---> Delete all the blood-stats.
	if err := initializers.DB.Delete(&bloodStatIds); err != nil {
		return errors.New("blood-stats cannot be deleted")
	}

	//----> Send back response.
	return nil
}

func (bloodStat *BloodStat) EditBloodStatById(id string, userAuth utils.UserAuth) error {
	//----> retrieve blood-stat and check for error.
	if _, err := getOneBloodStat(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Edit the blood-stat
	if err := initializers.DB.Model(&bloodStat).Updates(bloodStat).Error; err != nil {
		return errors.New("failed to update blood stat from database")
	}

	//----> Send back the response.
	return nil
}

func (b *BloodStat) GetBloodStatById(id string, userAuth utils.UserAuth) (BloodStat, error) {
	//----> Retrieve the blood-stat from database.
	bloodStat, err := getOneBloodStat(id, userAuth)

	//----> Check for error.
	if err != nil {
		return BloodStat{}, errors.New(err.Error())
	}

	//----> send back the response.
	return bloodStat, nil
}

func (bloodStat *BloodStat) GetBloodStatByUserId(userId string) (BloodStat, error){
	//----> Retrieve the blood-stat by user-id.
	if err := initializers.DB.First(&bloodStat, BloodStat{UserID: userId}); err != nil {
		return BloodStat{}, errors.New("the blood-stat for this user cannot be retrieved")
	}

	//----> Send back the response
	return *bloodStat, nil
}

func (b *BloodStat) GetAllBloodStat() ([]BloodStat, error) {
	var bloodStats []BloodStat //----> Declare the variable.

	//----> Retrieve all the blood-stats from database.
	if err := initializers.DB.Find(&bloodStats).Error; err != nil {
		return []BloodStat{}, errors.New("failed to retrieve blood stats")
	}

	//----> Send back the response.
	return bloodStats, nil
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

func getAllBloodStatIds(bloodStats []BloodStat) []BloodStat {
	bloodStatIds := []BloodStat{}

	//----> Collect all the blood-stat ids.
	for _, bloodStat := range bloodStats{
		bloodStat := BloodStat{ID: bloodStat.ID}

		bloodStatIds = append(bloodStatIds, bloodStat)
	}

	//----> send back the result
	return bloodStatIds
}
