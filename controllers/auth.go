package controllers

import (
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"

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
	token, err := login.Login()

	//----> Check the error.
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Store the token in the cookie.
	middlewares.SetCookieHandler(c, token)

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(token)
}

func LogoutController(c *fiber.Ctx) error {
	//----> Remove the cookie.
	middlewares.DeleteCookieHandler(c)

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successfully", "status": "success"})
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
