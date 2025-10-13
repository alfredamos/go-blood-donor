package models

import (
	"errors"
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Valid   Status = "Valid"
	Invalid Status = "Invalid"
)

type QueryConditions struct {
	Status      Status
	UserID      string
	accessToken string
}

type Token struct {
	ID           string `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt  `gorm:"index"`
	AccessToken  string          `json:"accessToken" gorm:"unique;type:varchar(750)"`
	RefreshToken string          `json:"refreshToken" gorm:"unique;type:varchar(750)"`
	TokenType    utils.TokenType `json:"tokenType"`
	Expired      bool            `json:"expired"`
	Revoked      bool            `json:"revoked"`
	Status       Status          `json:"status"`
	UserID       string          `json:"userId"`
}

// BeforeCreate These functions are called before creating any Post
func (token *Token) BeforeCreate(_ *gorm.DB) (err error) {
	token.ID = uuid.New().String()
	return
}

func (token *Token) CreateToken() (err error) {
	//----> Save the new token in its database.
	if err := initializers.DB.Create(&token).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response.
	return nil
}

func (*Token) DeleteInvalidTokensByUserId(userId string) error {
	//----> Retrieve all invalid tokens.
	queryConditions := QueryConditions{UserID: userId, Status: Invalid}
	return deleteInvalidTokens(queryConditions)
}

func (*Token) DeleteAllInvalidTokens() error {
	//----> Retrieve all invalid tokens.
	queryConditions := QueryConditions{Status: Invalid}
	return deleteInvalidTokens(queryConditions)
}

func (token *Token) FindTokenByAccessToken() (Token, error) {
	//----> Retrieve one token object by access token.
	queryConditions := QueryConditions{accessToken: token.AccessToken}
	if err := initializers.DB.Where(&queryConditions).First(&token).Error; err != nil {
		return Token{}, errors.New(err.Error())
	}

	//----> Send back result.
	return *token, nil
}

func (*Token) FindAllValidTokensByUserId(userId string) ([]Token, error) {
	//----> Retrieve all valid tokens.
	queryConditions := QueryConditions{UserID: userId, Status: Valid}
	tokens, err := findValidOrInvalidTokens(queryConditions)

	//----> Check for error in retrieving valid tokens.
	if err != nil {
		return nil, err
	}

	//----> Send back results.
	return tokens, nil
}

func (token *Token) RevokeAllValidTokensByUserId(userId string) error {
	//----> Retrieve all valid tokens.
	tokens, err := token.FindAllValidTokensByUserId(userId)

	//----> Check for errors in retrieving valid tokens.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Revoke all valid tokens
	if err := revokeValidTokens(tokens); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil

}

func MakeNewToken(userId string, token Token) Token {
	return Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expired:      false,
		Revoked:      false,
		TokenType:    utils.Bearer,
		Status:       Valid,
		UserID:       userId,
	}
}

func sliceOfTokenIds(tokens []Token) (ids []string) {
	//----> Make slice of token ids.
	for _, token := range tokens {
		ids = append(ids, token.ID)
	}

	//----> Send back the result.
	return ids
}

func revokeValidTokens(tokens []Token) error {
	revokedTokens := make([]Token, 0) //----> Initialize slice of token.

	//----> Revoke all tokens.
	for _, token := range tokens {
		token.Expired = true   //----> Token has expired.
		token.Revoked = true   //----> Token is revoked.
		token.Status = Invalid //----> Token is invalid.

		revokedTokens = append(revokedTokens, token) //----> Update slice of revoked tokens.
	}

	//----> Save all revoked tokens in the database.
	if err := saveAll(revokedTokens); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil
}

func findValidOrInvalidTokens(queryConditions QueryConditions) ([]Token, error) {
	//----> Initialize tokens.
	tokens := new([]Token)
	//----> Retrieve valid or invalid tokens from database.
	if err := initializers.DB.Where(&queryConditions).Find(&tokens).Error; err != nil {
		return []Token{}, errors.New(err.Error())
	}

	//----> Send back results.
	return *tokens, nil
}

func deleteInvalidTokens(queryConditions QueryConditions) error {
	invalidTokens, err := findValidOrInvalidTokens(queryConditions)
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Collect ids of invalid tokens in a slice.
	tokenIds := sliceOfTokenIds(invalidTokens)

	//----> Delete all invalid tokens from token database.
	if err := initializers.DB.Delete(&tokenIds).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil
}

func saveAll(tokens []Token) error {
	tx := initializers.DB.Begin()
	if tx.Error != nil {
		return errors.New("error at the onset of saving all tokens")
	}

	for _, token := range tokens {
		err := tx.Model(&Token{}).Where("id = ?", token.ID).Updates(token).Error

		//----> Check for error
		if err != nil {
			tx.Rollback()
			// Handle error
			return errors.New("error updating token")
		}
	}

	err := tx.Commit().Error
	if err != nil {
		// Handle error
		return errors.New("unable to commit token to database")
	}
	return nil
}
