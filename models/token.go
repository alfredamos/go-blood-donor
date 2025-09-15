package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Token struct{
	ID           string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	AccessToken string `json:"accessToken" gorm:"unique"`
	RefreshToken string `json:"refreshToken" gorm:"unique"`
	TokenType utils.TokenType `json:"tokenType"`
	Expired bool `json:"expired"`
	Revoked bool `json:"revoked"`
	UserID      string         `json:"userId"`
}

// BeforeCreate These functions are called before creating any Post
func (token *Token) BeforeCreate(_ *gorm.DB) (err error) {
	token.ID = uuid.New().String()
	return 
}

func (token *Token) findAllValidTokensByUser(userId string)([]Token, error) {
	tokens := new([]Token)

	//----> Get all tokens that match the specified criteria from the database.
	if err := initializers.DB.Find(&tokens, Token{UserID: userId, Expired: false, Revoked: false}); err != nil {
		return []Token{}, errors.New("tokens that match specified criteria are not available")
	}

	//----> Send back the results.
	return *tokens, nil
}

func (t *Token) getTokenByAccessToken(accessToken string)(Token, error){
	token := new(Token)

	//----> Retrieve the blood stat.
	if err := initializers.DB.Where(&Token{AccessToken: accessToken}).First(&token).Error; err != nil {
		return Token{}, errors.New(err.Error())
	}

	//----> Send back the results.
	return *token, nil
}