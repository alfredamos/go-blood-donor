package controllers

import (
	"errors"
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"
	"go-donor-list-backend/utils"

	"github.com/gofiber/fiber/v2"
)


func ChangePasswordController(c *fiber.Ctx) error {
	changePassword := new(models.ChangePasswordRequest)

	//----> Get the request body from context.
	if err := c.BodyParser(&changePassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Change the credentials in the database.
	if err := changePassword.ChangePassword(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})

	}

	//----> Send back the response.
	c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Change password successfully", "status": "success"})

	return nil
}

func GetCurrentUserController(c *fiber.Ctx) error {
	user := new(models.User)
	//----> Get the current user-id from context.
	email := middlewares.GetUserEmailFromContext(c)

	//----> Get the current user info from database.
	currentUser, err := user.GetCurrentUser(email)

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(200).JSON(fiber.Map{"currentUser": currentUser})
}

func EditProfileController(c *fiber.Ctx) error {
	editProfile := new(models.EditProfileRequest)

	//----> Get the request body from context.
	if err := c.BodyParser(&editProfile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})

	}

	//----> Edit user profile in the database.
	if err := editProfile.EditProfile(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User profile edited successfully", "status": "success"})

}

func LoginController(c *fiber.Ctx) error {
	login := new(models.LoginRequest)

	//----> Get the request body from context.
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Login the user.
	accessToken, refreshToken, err := login.Login()

	//----> Check the error.
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Store both access-token and refresh-token in the cookie.
	middlewares.SetCookieHandler(c, string(utils.AccessTokenPath), string(utils.AccessToken), accessToken)
	middlewares.SetCookieHandler(c, string(utils.RefreshTokenPath), string(utils.RefreshToken), refreshToken)

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(accessToken)

}

func LogoutController(c *fiber.Ctx) error {
	//----> Fetch the access token.
	accessToken := middlewares.GetCookieHandler(c, "accessToken")
	
	//---> Validate token
	isValid, err := middlewares.IsValid(c, accessToken)
	
	//----> Check for error.
	if err != nil{
		return errors.New("error validating token")
	}
	
	//----> Check for valid token
	if !isValid{
		return errors.New("invalid or expired token")
	}

	//----> Invalidate the token in the database.
	models.Logout(accessToken)

	//----> Remove the cookie.
	middlewares.DeleteCookieHandler(c, "accessToken")

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successfully", "status": "success"})
}

func RefreshTokenController(c *fiber.Ctx)error{
	//----> Get refreshToken.
	refreshToken := middlewares.GetCookieHandler(c, string(utils.RefreshToken))
	
	//---> Validate token
	isValid, err := middlewares.IsValid(c, refreshToken)
	
	//----> Check for error.
	if err != nil{
		return errors.New("error validating token")
	}

	//----> Check for valid token
	if !isValid{
		return errors.New("invalid or expired token")
	}

	//----> Get the user.

	userDetail := middlewares.GetUser_Id_Email_Name_Role(c)
	
	newAccessToken, newRefreshToken, err := models.RefreshToken(userDetail)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Set the access-token and refresh-token on context cookies.
	middlewares.SetCookieHandler(c, string(utils.AccessTokenPath),string(utils.AccessToken), newAccessToken)
	middlewares.SetCookieHandler(c, string(utils.RefreshTokenPath), string(utils.RefreshToken), newRefreshToken)

	//----> Send back response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "token has been refreshed successfully"})
}

func SignupController(c *fiber.Ctx) error {
	signup := new(models.SignupRequest)

	//----> Get the signup request from context.
	if err := c.BodyParser(&signup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Signup the new user.
	if err := signup.Signup(); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Signup successfully", "status": "success"})
}
