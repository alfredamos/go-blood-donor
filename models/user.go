package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/utils"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"name" binding:"required"`
	Address      utils.Address        `json:"address" gorm:"embedded"`
	Email        string         `json:"email" binding:"required" gorm:"unique"`
	Image        string         `json:"image" binding:"required"`
	Phone        string         `json:"phone" binding:"required"`
	Password     string         `json:"-" binding:"required"`
	Gender       utils.Gender   `json:"gender" gorm:"default:Male"`
	DateOfBirth  time.Time      `json:"dateOfBirth" binding:"required"`
	Age          int            `json:"age"`
	Role         utils.Role     `json:"role" gorm:"default:'Customer'"`
	Vitals       []Vital        `json:"vitals" gorm:"foreignKey:UserID"`
	DonorDetails []DonorDetail  `json:"donorDetails" gorm:"foreignKey:UserID"`
	BloodStat    BloodStat      `json:"bloodStat" gorm:"foreignKey:UserID"`
}

// BeforeCreate These functions are called before creating any Post
func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return
}

func (d *User) DeleteUserById(id string, userAuth utils.UserAuth) error {
	//----> Check for existence of user.
	if _, err := getOneUser(id, userAuth); err != nil {
		return errors.New(err.Error())
	}

	//----> Delete the user.
	if err := initializers.DB.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return errors.New("failed to delete user")
	}

	//----> Send back the response.
	return nil

}

func (d *User) GetUserById(id string, userAuth utils.UserAuth) (User, error) {
	//----> Retrieve the user from database.
	user, err := getOneUser(id, userAuth)

	//----> Check for error.
	if err != nil {
		return User{}, errors.New(err.Error())
	}

	//----> Send back the response.
	return user, nil
}

func (d *User) GetAllUsers() ([]User, error) {
	var users []User //----> Declare the variable.

	//----> Retrieve the users from the database.
	if err := initializers.DB.Find(&users).Error; err != nil {
		return users, errors.New("failed to retrieve users from database")
	}

	//----> Send back the response.
	return users, nil

}

func getOneUser(id string, userAuth utils.UserAuth) (User, error) {
	user := User{} //----> User variable.

	//----> Retrieve the user with the given id from the database.
	if err := initializers.DB.First(&user, "id = ?", id).Error; err != nil {
		return User{}, errors.New("failed to find user with id " + id + " in database")
	}

	//----> Check for ownership and admin privilege.
	if err := utils.CheckForOwnership(userAuth.UserId, user.ID, userAuth.IsAdmin); err != nil{
		return User{}, errors.New("you are not permitted to view or perform any action on this resource")
	}

	//----> Send back the response.
	return user, nil
}
