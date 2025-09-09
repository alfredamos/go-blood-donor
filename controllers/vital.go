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
	
	//----> Store the user-id in vital.
	vital.UserID = userId

	//----> store the newly created vital in the database.
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

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Delete the blood-stat with given id from the database.
	if err := vital.DeleteVitalById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vital has been deleted successfully!", "status": "success"})
}

func DeleteAllVitalsController(c *fiber.Ctx) error{
	vital := new(models.Vital)

	//----> Delete all vitals from the database.
	if err := vital.DeleteAllVitals(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "All vitals have been deleted successfully!", "status": "success"})
}
func DeleteAllVitalsByUserIdController(c *fiber.Ctx) error{
	vital := new(models.Vital)
	
	//----> Get user-id from params.
	userId := c.Params("userId")

	//----> Delete all vitals from the database.
	if err := vital.DeleteAllVitalsByUserId(userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "All vitals have been deleted successfully!", "status": "success"})
}

func EditVitalByIdController(c *fiber.Ctx) error{
	vital := new(models.Vital)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the edited blood-stat payload from the context.
	if err := c.BodyParser(&vital); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})	
	}

	//----> Store the user-id in vital.
	vital.UserID = userAuth.UserId

	//----> Update the blood-stat with given id from the database.
	if err := vital.EditVitalById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Vital has been deleted successfully!", "status": "success"})
}

func GetVitalByIdController(c *fiber.Ctx) error {
	vital := new(models.Vital)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the blood-stat with given id from the database.
	fetchedVital, err := vital.GetVitalById(id, userAuth)
	
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

func GetAllVitalsByUserIdController(c *fiber.Ctx)error{
	vital := new(models.Vital)

	//----> Get the user-id from params.
	userId := c.Params("userId")

	//----> Retrieve vitals by user-id.
	vitals, err := vital.GetAllVitalsByUserId(userId)

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": vitals, "status": "success"})
}
