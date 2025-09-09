package responses

import (
	"go-donor-list-backend/utils"
	"time"
)

type UserResponse struct {
	ID					string			 `json:"id"`
	Name        string       `json:"name" binding:"required"`
	Address     utils.Address      `json:"address"`
	Email       string       `json:"email" binding:"required"`
	Image       string       `json:"image" binding:"required"`
	Phone       string       `json:"phone" binding:"required"`
	Gender      utils.Gender `json:"gender"`
	DateOfBirth time.Time    `json:"dateOfBirth" binding:"required"`
	Age         int          `json:"age"`
	Role        utils.Role   `json:"role"`
}