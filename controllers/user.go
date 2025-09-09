package controllers

import (
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"

	"github.com/gofiber/fiber/v2"
)

func DeleteUserByIdController(c *fiber.Ctx) error{
	user := new(models.User)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id of user from params.
	id := c.Params("id")

	//----> Get the user with the given id from database.
	if err := user.DeleteUserById(id, userAuth); err !=nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user is deleted successfullY!", "status": "success"})
}

func GetAllUsersController(c *fiber.Ctx) error{
	user := new(models.User)

	//----> Get the user with the given id from database.
	users, err := user.GetAllUsers()
	if err !=nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": users, "status": "success"})
}

func GetUserByIdController(c *fiber.Ctx) error{
	user := new(models.User)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id of user from params.
	id := c.Params("id")

	//----> Get the user with the given id from database.
	fetchedUser, err := user.GetUserById(id, userAuth)
	if err !=nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fetchedUser, "status": "success"})
}
