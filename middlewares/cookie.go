package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookieHandler(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
	})

}

func GetCookieHandler(c *fiber.Ctx) string {
	return c.Cookies("token")
}

func DeleteCookieHandler(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",                               // Clear the value
		Expires:  time.Now().Add(-3 * time.Second), // Set an expired time for older browsers (optional but recommended)
		MaxAge:   -1,                               // Set MaxAge to a negative value to delete the cookie immediately
		HTTPOnly: true,                             // Should match the original cookie's HttpOnly setting
		Path:     "/",                              // Should match the original cookie's Path setting
		Domain:   "localhost",                      // Should match the original cookie's Domain setting (or leave empty for host cookie)
		Secure:   false,
	})
}
