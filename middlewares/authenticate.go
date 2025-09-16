package middlewares

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(name, email, userId, role string)(string, error){
	expireTime := time.Now().Add(time.Minute * 15).Unix()

	//----> Generate access-token
	accessToken, err := generateToken(name, email, userId, role, expireTime)

	//----> Check for error.
	if err != nil {
		return "", errors.New("error generating access-token")
	}

	//----> send back the result.
	return accessToken, nil
}

func GenerateRefreshToken(name, email, userId, role string)(string, error){
	expireTime := time.Now().Add(time.Hour * 24 * 7).Unix()

	//----> Generate access-token
	accessToken, err := generateToken(name, email, userId, role, expireTime)

	//----> Check for error.
	if err != nil {
		return "", errors.New("error generating access-token")
	}

	//----> send back the result.
	return accessToken, nil
}

func generateToken(name string, email string, userId string, role string, expireAt int64) (string, error) {
	secretKey := os.Getenv("JWT_TOKEN_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"name": name, "email": email, "userId": userId, "role": role, "expiresAt": expireAt})
	return token.SignedString([]byte(secretKey))
}

func VerifyTokenJwt(c *fiber.Ctx) error {
	//----> Get token from cookie.
	token := GetCookieHandler(c, "accessToken")

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
