package controllers

import (
	"errors"
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateBloodStatController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)

	//----> Get the user id.
	userId := middlewares.GetUserIdFromContext(c)

	//----> Get the bloodStat payload from context.
	if err := c.BodyParser(&bloodStat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please provide all values!", "status": "fail"})
	}

	//----> Store the user-id in blood-stat.
	bloodStat.UserID = userId

	//----> Store the newly created blood stats in the database.
	newBloodStat, err := bloodStat.CreateBloodStat()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": newBloodStat, "status": "success"})
}

func DeleteBloodStatByIdController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Delete the blood-stat with given id from the database.
	if err := bloodStat.DeleteBloodStatById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "BloodStat has been deleted successfully!", "status": "success"})
}

func DeleteBloodStatByUserIdController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)
	//----> Get the id from context params.
	userId := c.Params("userId")

	//----> Delete the blood-stat with given id from the database.
	if err := bloodStat.DeleteBloodStatByUserId(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "BloodStat has been deleted successfully!", "status": "success"})
}

func DeleteAllBloodStatController(c *fiber.Ctx) error {
	bloodStat := new(models.BloodStat)

	//----> Delete all the blood-stats.
	if err := bloodStat.DeleteAllBloodStat(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "All blood-stats have been deleted successfully!", "status": "success"})
}

func EditBloodStatByIdController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the edited blood-stat payload from the context.
	if err := c.BodyParser(&bloodStat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})	
	}

	//----> Store the user-id in blood-stat.
	bloodStat.UserID = userAuth.UserId

	//----> Update the blood-stat with given id from the database.
	if err := bloodStat.EditBloodStatById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "BloodStat has been deleted successfully!", "status": "success"})
}

func GetBloodStatByIdController(c *fiber.Ctx) error {
	bloodStat := new(models.BloodStat)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the blood-stat with given id from the database.
	fetchedBloodStat, err := bloodStat.GetBloodStatById(id, userAuth)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fetchedBloodStat, "status": "success"})
}

func GetBloodStatByUserIdController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)

	//----> Get the user-id from params.
	userId := c.Params("userId")

	//----> Retrieve the blood-stat associated with this user.
	userBloodStat, err := bloodStat.GetBloodStatByUserId(userId)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": userBloodStat, "status": "success"})
}

func GetAllBloodStatsController(c *fiber.Ctx) error {
	bloodStat := new(models.BloodStat)

	//----> Get all the blood-stats from database.
	allBloodStats, err := bloodStat.GetAllBloodStat()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": allBloodStats, "status": "success"})
}

