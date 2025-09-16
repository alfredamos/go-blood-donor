package utils

type TokenName string

const (
	AccessToken TokenName = "accessToken"
	RefreshToken TokenName = "refreshToken"
)

type TokenPath string

const (
	AccessTokenPath TokenPath = "/"
	RefreshTokenPath TokenPath = "/api/auth/refresh"
)
