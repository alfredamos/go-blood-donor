package models

import (
	"errors"
	"fmt"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/middlewares"
	_ "net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordRequest struct {
	Email           string `json:"email" binding:"required;email"`
	ConfirmPassword string `json:"confirmPassword"`
	NewPassword     string `json:"newPassword"`
	OldPassword     string `json:"oldPassword"`
}

func (req *ChangePasswordRequest) ChangePassword() error {
	email := req.Email
	confirmPassword := req.ConfirmPassword
	newPassword := req.NewPassword
	oldPassword := req.OldPassword

	user := new(User)

	//----> Check for match between newPassword and confirmPassword.
	if isMatch := passwordMatch(newPassword, confirmPassword); !isMatch {
		return errors.New("password does not match")
	}

	//----> Check for existence of user.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Check for match between oldPassword and the one in the database.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Hash the newPassword.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	//----> Check for error.
	if err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Save the hashedPassword in the database.
	user.Password = string(hashedPassword)
	initializers.DB.Save(&user)

	//----> Send back the response.
	return nil

}

type EditProfileRequest struct {
	Name        string    `json:"name" binding:"required"`
	Address     Address   `json:"address" gorm:"embedded"`
	Email       string    `json:"email" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Gender      Gender    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	Age         int       `json:"age"`
	Role        Role      `json:"role" gorm:"default:'Customer'"`
	Password    string    `json:"password" binding:"required"`
}

func (req *EditProfileRequest) EditProfile() error {
	email := req.Email
	password := req.Password

	user := new(User)

	//----> Check for existence of user.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Check for password correctness.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Update the user in the database.
	updatedUser := editProfileRequestToUser(req)
	fmt.Print("Edited user, updatedUser : ", updatedUser)
	if err := initializers.DB.Model(&user).Updates(&updatedUser).Error; err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Send back the response.
	return nil
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required;email"`
	Password string `json:"password" binding:"required"`
}

func (req *LoginRequest) Login() (string, error) {
	email := req.Email
	password := req.Password

	user := new(User)

	//----> Check for existence of user.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials ")
	}

	//----> Check for match between oldPassword and the one in the database.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials ")
	}

	//----> Generate token.
	token, err := middlewares.GenerateToken(user.Name, user.Email, user.ID, string(user.Role))

	//----> Check for error.
	if err != nil {
		return "", errors.New("expired or invalid credentials")
	}

	//----> Send back the response.
	return token, nil
}

type SignupRequest struct {
	Name            string    `json:"name" binding:"required"`
	Address         Address   `json:"address" gorm:"embedded"`
	Email           string    `json:"email" binding:"required"`
	Phone           string    `json:"phone" binding:"required"`
	Image           string    `json:"image" binding:"required"`
	Gender          Gender    `json:"gender" binding:"required"`
	DateOfBirth     time.Time `json:"dateOfBirth" binding:"required"`
	Age             int       `json:"age"`
	Password        string    `json:"password" binding:"required"`
	ConfirmPassword string    `json:"confirmPassword" binding:"required"`
}

func (user *User) GetCurrentUser(email string) (User, error) {
	//----> Get current user from the database.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return User{}, errors.New("invalid credentials ")
	}

	//----> Send back the response.
	return *user, nil
}

func (req *SignupRequest) Signup() error {
	fmt.Println("NewUser : ", req)
	email := req.Email
	confirmPassword := req.ConfirmPassword
	password := req.Password

	user := new(User)

	//----> Check for match between password and confirmPassword.
	if isMatch := passwordMatch(password, confirmPassword); !isMatch {
		return errors.New("password does not match")
	}

	//----> Check for existence of user.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return errors.New("user already exists")
	}

	//----> Hash password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	//----> Check for error.
	if err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Mapped signupRequest to user.
	newUser := signupRequestToUser(req, string(hashedPassword))

	//----> save the new user in the database.
	if err := initializers.DB.Create(&newUser).Error; err != nil {
		return errors.New("invalid credentials ")
	}

	//----> send back the response.
	return nil
}

func passwordMatch(password, confirmPassword string) bool {
	return password == confirmPassword
}

func signupRequestToUser(req *SignupRequest, hashedPassword string) User {
	return User{
		Address:     req.Address,
		Email:       req.Email,
		Name:        req.Name,
		Password:    hashedPassword,
		Phone:       req.Phone,
		Image:       req.Image,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Age:         calculateAge(req.DateOfBirth),
		Role:        Customer,
	}
}

func editProfileRequestToUser(req *EditProfileRequest) User {
	return User{
		Address:     req.Address,
		Name:        req.Name,
		Phone:       req.Phone,
		Image:       req.Image,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Age:         calculateAge(req.DateOfBirth),
	}
}

func calculateAge(dateOfBirth time.Time)int {
	exactAge := time.Now().Year() - dateOfBirth.Year()
	currentMonth := time.Now().Month()
	currentDayOfMonth := time.Now().Day()

	monthOfBirth := dateOfBirth.Month()
	dayOfBirth := dateOfBirth.Day()

	if ((monthOfBirth == currentMonth) && (dayOfBirth >= currentDayOfMonth)) || (currentMonth >= monthOfBirth) {
		return exactAge
	}

	return exactAge - 1
}
