package controllers

import (
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateVitalController(c *fiber.Ctx) error{
	vital := new(models.Vital)

	//----> Get the user id.
	userId := middlewares.GetUserIdFromContext(c)

	//----> Get the vital payload from context.
	if err := c.BodyParser(&vital); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please provide all values!", "status": "fail"})
	}

	//----> Store the user blood-stat in the database.
	vital.UserID = userId
	newVital, err := vital.CreateVital()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": newVital, "status": "success"})
}

func DeleteVitalByIdController(c *fiber.Ctx) error{
	vital := new(models.Vital)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Delete the blood-stat with given id from the database.
	if err := vital.DeleteVitalById(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vital has been deleted successfully!", "status": "success"})
}

func EditVitalByIdController(c *fiber.Ctx) error{
	vital := new(models.Vital)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the edited blood-stat payload from the context.
	if err := c.BodyParser(&vital); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})	
	}

	//----> Update the blood-stat with given id from the database.
	if err := vital.EditVitalById(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vital has been deleted successfully!", "status": "success"})
}

func GetVitalByIdController(c *fiber.Ctx) error {
	vital := new(models.Vital)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the blood-stat with given id from the database.
	fetchedVital, err := vital.GetVitalById(id)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fetchedVital, "status": "success"})
}

func GetAllVitalsController(c *fiber.Ctx) error {
	vital := new(models.Vital)

	//----> Get all the blood-stats from database.
	allVitals, err := vital.GetAllVitals()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": allVitals, "status": "success"})
}

