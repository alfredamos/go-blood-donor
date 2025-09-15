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
	AccessToken string `json:"accessToken" gorm:"unique;type:varchar(750)"`
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

func FindAllValidTokensByUser(userId string)([]Token, error) {
	tokens := new([]Token)

	//----> Get all tokens that match the specified criteria from the database.
	if err := initializers.DB.Find(&tokens, Token{UserID: userId, Expired: false, Revoked: false}); err != nil {
		return []Token{}, errors.New("tokens that match specified criteria are not available")
	}

	//----> Send back the results.
	return *tokens, nil
}

func GetTokenByAccessToken(accessToken string)(Token, error){
	token := new(Token)

	//----> Retrieve the blood stat.
	if err := initializers.DB.Where(&Token{AccessToken: accessToken}).First(&token).Error; err != nil {
		return Token{}, errors.New(err.Error())
	}

	//----> Send back the results.
	return *token, nil
}

func (t *Token) DeleteTokensByUserId(userId string) error{
	tokens := new(Token)

	//----> Get all tokens.
	if err := initializers.DB.Find(&tokens, Token{UserID: userId}).Error; err != nil {
		return errors.New("tokens cannot be retrieve from database")
	}

	//----> Delete tokens for a particular user by his/her id.
	if err := initializers.DB.Where(&Token{UserID: userId}).Find(&Token{}).Error; err != nil{
		return errors.New("error deleting tokens")
	}

	//----> send back response.
	return nil
}

func (t *Token) DeleteAllTokens() error{
	tokens := []Token{}

	//----> Get all tokens.
	if err := initializers.DB.Find(&tokens).Error; err != nil {
		return errors.New("tokens cannot be retrieve from database")
	}

	tokenIds := getAllTokensId(tokens)

	//----> Delete all tokens.
	if err := initializers.DB.Delete(tokenIds).Error; err != nil {
		return errors.New("error deleting tokens")
	}

	//----> send back response.
	return nil
}

func getAllTokensId(tokens []Token) []Token {
	tokenIds := make([]Token,0) 

	//----> Get all the token ids.
	for _, token := range tokens {
		tokenIds = append(tokenIds, Token{ID: token.ID})
	}

	//----> Send back the results.
	return tokenIds
}