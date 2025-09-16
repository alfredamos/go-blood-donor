package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/utils"
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
	Address     utils.Address   `json:"address" gorm:"embedded"`
	Email       string    `json:"email" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Phone       string    `json:"phone" binding:"required"`
	Gender      utils.Gender    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" binding:"required"`
	Age         int       `json:"age"`
	Role        utils.Role      `json:"role" gorm:"default:'Customer'"`
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
	
	if err := initializers.DB.Model(&user).Updates(updatedUser).Error; err != nil {
		return errors.New("invalid credentials ")
	}

	//----> Send back the response.
	return nil
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required;email"`
	Password string `json:"password" binding:"required"`
}

func (req *LoginRequest) Login() (string, string, error) {
	email := req.Email
	password := req.Password

	user := new(User)
	newToken := new(Token)

	//----> Check for existence of user.
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", errors.New("invalid credentials ")
	}

	//----> Check for match between oldPassword and the one in the database.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials ")
	}

	//----> Revoke tokens.
	
	if err := revokeAllUserTokens(user.ID); err != nil {
		return "", "", errors.New(err.Error())
	}

	//----> Generate access token.
	accessToken, err := middlewares.GenerateAccessToken(user.Name, user.Email, user.ID, string(user.Role))

	//----> Check for error in generating.
	if err != nil {
		return "", "", errors.New("error generating access-token")
	}

	//----> Generate refresh-token
	refreshToken, err := middlewares.GenerateRefreshToken(user.Name, user.Email, user.ID, string(user.Role))

	//----> Check for error in generating refresh token.
	if err != nil {
		return "", "", errors.New("error generating refresh-token")		
	}
	//----> Store the token in the database.
	newToken.AccessToken = accessToken
	newToken.RefreshToken = refreshToken
	newToken.Expired = false
	newToken.Revoked = false
	newToken.Status = Valid
	newToken.UserID = user.ID
	newToken.TokenType = utils.Bearer

	if err := initializers.DB.Create(&newToken).Error; err != nil {
		return "", "", errors.New("error saving token in the database")
	}

	//----> Send back the response.
	return accessToken, refreshToken, nil
}

func Logout(accessToken string) error{
	//----> Get the current valid token.
  validToken, err := GetTokenByAccessToken(accessToken)

	//---> Invalidate token.
	validToken.Expired = true
	validToken.Revoked = true
	validToken.Status = Invalid

	//----> Save the updated token in the database.
	if err := initializers.DB.Save(validToken).Error; err != nil {
		return errors.New("error updating token")
	}

	//----> Check for error.
	if err != nil {
		return errors.New("error fetching token")
	}

	//----> Send back the response.
	return nil
}

type SignupRequest struct {
	Name            string    `json:"name" binding:"required"`
	Address         utils.Address   `json:"address" gorm:"embedded"`
	Email           string    `json:"email" binding:"required"`
	Phone           string    `json:"phone" binding:"required"`
	Image           string    `json:"image" binding:"required"`
	Gender          utils.Gender    `json:"gender" binding:"required"`
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
		Role:        utils.Role(utils.Customer),
	}
}

func RefreshToken(userDetail utils.UserDetail)(string, string, error) {
	editedToken := new(Token)
	
	//----> Get all valid tokens.
	validTokens, err := FindAllValidTokensByUser(userDetail.UserId)
	
	//----> Check for error.
	if err != nil {
		return "", "", errors.New("error retrieving valid tokens");
	}

	//----> Get the first token.
	validToken := validTokens[0]

	//---> Check token status.
  if validToken.Revoked && validToken.Expired{
    return "", "", errors.New("invalid or expired token");
  }

	//----> Revoke all tokens.
	if err := revokeAllUserTokens(userDetail.UserId); err != nil {
		return "", "", errors.New("error in revoking token")
	}

	//----> Generate access token.
	accessToken, err := middlewares.GenerateAccessToken(userDetail.Name, userDetail.Email, userDetail.UserId, string(userDetail.Role))

	//----> Check for error in generating.
	if err != nil {
		return "", "", errors.New("error generating access-token")
	}

	//----> Generate refresh-token
	refreshToken, err := middlewares.GenerateRefreshToken(userDetail.Name, userDetail.Email, userDetail.UserId, string(userDetail.Role))

	//----> Check for error in generating refresh token.
	if err != nil {
		return "", "", errors.New("error generating refresh-token")		
	}
	//----> Store the token in the database.
	editedToken.AccessToken = accessToken
	editedToken.RefreshToken = refreshToken
	editedToken.Expired = false
	editedToken.Revoked = false
	editedToken.Status = Valid
	editedToken.UserID = userDetail.UserId
	editedToken.TokenType = utils.Bearer

	if err := initializers.DB.Save(&editedToken).Error; err != nil {
		return "", "", errors.New("error saving token in the database")
	} 

	//----> send back response.
	return accessToken, refreshToken, nil
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
		//Role:        req.Role,
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

func revokeAllUserTokens(userId string) error{
	tokens := make([]Token,0)

	//----> Fetch all valid tokens.
	validTokens, err := FindAllValidTokensByUser(userId)
	
	// //----> Check for empty slice.
	if len(validTokens) == 0 {
		return nil
	}

	//----> Check for error.
	if err != nil {
		return errors.New("error fetching tokens")
	}

	//----> Collect tokens to revoke in a slice with expired and revoked as true.
	tokens = tokensToRevoke(userId, validTokens)

	//----> Store the updated tokens in the database.
	if err := saveAll(tokens); err != nil {
		return errors.New("error revoking tokens")
	}
	//----> Send back the response.
	return nil
}

func tokensToRevoke(userId string, validTokens []Token)[]Token{
	tokens := make([]Token, 0)

	for _, token := range validTokens{
		token.Expired = true //----> Token has expired.
		token.Revoked = true //----> Token has been revoked.
		token.Status = Invalid

		//----> Make new token.
		newToken := makeToken(userId, token)

		tokens = append(tokens, newToken)
	}

	//----> Send back the response.
	return tokens
}

func makeToken(userId string, token Token)Token{
	return Token {
		ID: token.ID,
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
		Revoked: token.Revoked,
		Status: token.Status,
		Expired: token.Expired,
		TokenType: token.TokenType,
		UserID: userId,
	}
}

func saveAll(tokens []Token)error{
	tx := initializers.DB.Begin()
	if tx.Error != nil {
		return errors.New("error at the onset of saving all tokens")
	}

	for _, token := range tokens {
		err := tx.Model(&Token{}).Where("id = ?", token.ID).Updates(token).Error
	
		//----> Check for error
		if err != nil {
			tx.Rollback()
			// Handle error
			return errors.New("error updating token")
		}
	}

	err := tx.Commit().Error
	if err != nil {
	// Handle error
	return errors.New("unable to commit token to database")
}
	return nil
}
