package controllers

import (
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

	//----> Delete the blood-stat with given id from the database.
	if err := bloodStat.DeleteBloodStatById(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "BloodStat has been deleted successfully!", "status": "success"})
}

func EditBloodStatByIdController(c *fiber.Ctx) error{
	bloodStat := new(models.BloodStat)

	//----> Get the user id.
	userId := middlewares.GetUserIdFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the edited blood-stat payload from the context.
	if err := c.BodyParser(&bloodStat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})	
	}

	//----> Store the user-id in blood-stat.
	bloodStat.UserID = userId

	//----> Update the blood-stat with given id from the database.
	if err := bloodStat.EditBloodStatById(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "BloodStat has been deleted successfully!", "status": "success"})
}

func GetBloodStatByIdController(c *fiber.Ctx) error {
	bloodStat := new(models.BloodStat)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the blood-stat with given id from the database.
	fetchedBloodStat, err := bloodStat.GetBloodStatById(id)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fetchedBloodStat, "status": "success"})
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

