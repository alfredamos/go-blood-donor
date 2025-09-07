package middlewares

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(name string, email string, userId string, role string) (string, error) {
	secretKey := os.Getenv("JWT_TOKEN_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": name, "email": email, "userId": userId, "role": role, "expiresAt": time.Now().Add(time.Hour * 2).Unix()})
	return token.SignedString([]byte(secretKey))
}

func VerifyTokenJwt(c *fiber.Ctx) error {
	//----> Get token from cookie.
	token := GetCookieHandler(c)

	//----> Validate token.
	parsedToken, err := validateToken(token)

	//----> Check for error.
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token!"})
	}

	//----> Get user claims.
	if err = getUserClaims(c, parsedToken); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token!"})
	}

	//----> User is authenticated.
	return c.Next()
}

func getUserClaims(c *fiber.Ctx, parsedToken jToken) error {
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		//----> Access claims
		name := claims["name"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)
		userId := claims["userId"]

		//----> Set the claims on gin context
		c.Locals("name", name)
		c.Set("name", name)

		c.Locals("email", email)
		c.Set("email", email)

		c.Locals("role", role)
		c.Set("role", role)

		//----> Convert user-id to string
		c.Locals("userId", userId)
		c.Set("userId", fmt.Sprintf("%v", userId))
		
		return nil
	}

	//----> User does not have claims.
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid token!"})

}

type jToken *jwt.Token

func validateToken(token string) (jToken, error) {
	secretKey := os.Getenv("JWT_TOKEN_SECRET")
	fmt.Println("In validateToken, secretKey : ", secretKey)
	//----> Parse token.
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		//----> Return the secret key for signing
		return []byte(secretKey), nil
	})

	//----> Check for error.
	if err != nil {
		return nil, errors.New("invalid credential")
	}

	//----> Check the validity of token.
	isValidToken := parsedToken.Valid

	//----> Check for error.
	if !isValidToken {
		return nil, errors.New("invalid credential")
	}

	//----> Send back the parsed token.
	return parsedToken, nil
}
