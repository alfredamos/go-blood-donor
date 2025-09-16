package middlewares

import (
	"errors"
	"fmt"
	"go-donor-list-backend/utils"

	"github.com/gofiber/fiber/v2"
)


func GetUserAuthFromContext(c *fiber.Ctx) utils.UserAuth{
	//----> Get user role from context.
	role := fmt.Sprintf("%v", c.Locals("role"))

	//----> Get the user-id from context.
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	
	isAdmin := role == "Admin"

	//----> Send back the role.
	return utils.UserAuth{IsAdmin: isAdmin, UserId: userId, Role: role}

}

func GetUser_Id_Email_Name_Role(c *fiber.Ctx)utils.UserDetail{
	//----> Get user name
	email := fmt.Sprintf("%v", c.Locals("email"))

	//----> Get user name
	name := fmt.Sprintf("%v", c.Locals("name"))

	//----> Get user role from context.
	role := fmt.Sprintf("%v", c.Locals("role"))

	//----> Get the user-id from context.
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	
	//----> Send back user-detail
	return utils.UserDetail{Email: email, Name: name,Role: role, UserId: userId}
}

func GetUserEmailFromContext(c *fiber.Ctx)string{
	//----> Get user-id from context.
	email := c.Locals("email")

	//----> Send back the user-id.
	return fmt.Sprintf("%v", email)
}

func GetUserIdFromContext(c *fiber.Ctx) string {
	//----> Get user-id from context.
	userId := c.Locals("userId")

	//----> Send back the user-id.
	return fmt.Sprintf("%v", userId)
}

func IsValid(c *fiber.Ctx, token string)(bool, error) {
	parsedToken, err := validateToken(token)

	//----> Check for errors.
	if err != nil {
		return false, errors.New("Invalid or expired token")
	}

	//----> Get-user claims
	if err := getUserClaims(c, parsedToken); err != nil {
		return false, errors.New("user claims are retrievable because of invalid or expired token")
	}

	

	//---> Send back the results.
	return parsedToken.Valid, nil
}
