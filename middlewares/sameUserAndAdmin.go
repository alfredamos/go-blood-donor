package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SameUserAndAdminMiddleware(c *fiber.Ctx) error{
	//----> Get the user-id from param.
	userIdFromContext := c.Params("userId")
	fmt.Println("In same-user-and-admin, c : ", c)
	//----> Get user role from context.
	userAuth := GetUserAuthFromContext(c)

	//----> Check for same user.
	isUserSame := IsSameUser(userAuth.UserId, userIdFromContext)

	//----> Check for same user and admin privilege.
	if !isUserSame && !userAuth.IsAdmin {
		//----> Invalid user.
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "You are not permitted to access this page!", "statusCode": http.StatusForbidden})
	}

	//----> Same user and admin are allowed to gain access.
	return c.Next()
}

func IsSameUser(userId1, userId2 string) bool {
	return userId1 == userId2
}
