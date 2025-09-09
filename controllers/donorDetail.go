package controllers

import (
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateDonorDetailController(c *fiber.Ctx) error{
	donorDetail := new(models.DonorDetail)

	//----> Get the user id.
	userId := middlewares.GetUserIdFromContext(c)

	//----> Get the donorDetail payload from context.
	if err := c.BodyParser(&donorDetail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Please provide all values!", "status": "fail"})
	}

	//----> Store the user-id in donor-details.
	donorDetail.UserID = userId

	//----> Store the newly created donor-detail in the database.
	newDonorDetail, err := donorDetail.CreateDonorDetail()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": newDonorDetail, "status": "success"})
}

func DeleteDonorDetailByIdController(c *fiber.Ctx) error{
	donorDetail := new(models.DonorDetail)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Delete the blood-stat with given id from the database.
	if err := donorDetail.DeleteDonorDetailById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "DonorDetail has been deleted successfully!", "status": "success"})
}

func EditDonorDetailByIdController(c *fiber.Ctx) error{
	donorDetail := new(models.DonorDetail)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the edited blood-stat payload from the context.
	if err := c.BodyParser(&donorDetail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "status": "fail"})	
	}

	//----> Store the user-id in donor-details.
	donorDetail.UserID = userAuth.UserId

	//----> Update the blood-stat with given id from the database.
	if err := donorDetail.EditDonorDetailById(id, userAuth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "DonorDetail has been deleted successfully!", "status": "success"})
}

func GetDonorDetailByIdController(c *fiber.Ctx) error {
	donorDetail := new(models.DonorDetail)

	//----> Get the user-auth.
	userAuth := middlewares.GetUserAuthFromContext(c)

	//----> Get the id from context params.
	id := c.Params("id")

	//----> Get the blood-stat with given id from the database.
	fetchedDonorDetail, err := donorDetail.GetDonorDetailByID(id, userAuth)
	
	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": fetchedDonorDetail, "status": "success"})
}

func GetAllDonorDetailsController(c *fiber.Ctx) error {
	donorDetail := new(models.DonorDetail)

	//----> Get all the blood-stats from database.
	allDonorDetails, err := donorDetail.GetAllDonorDetails()

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": allDonorDetails, "status": "success"})
}

func GetAllDonorDetailsByUserIdController(c *fiber.Ctx)error {
	donorDetail := new(models.DonorDetail)

	//----> Get the user-id from params.
	userId := c.Params("userId")

	//----> Retrieve all the donor-details associated with this user-id
	donorDetails, err := donorDetail.GetAllDonorDetailsByUserId(userId)

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "status": "fail"})
	}

	//----> Send back the response.
	return  c.Status(fiber.StatusOK).JSON(fiber.Map{"message": donorDetails, "status": "success"})
}



