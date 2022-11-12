package auth

import (
	"dormitory-system/src/model"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const RefreshTokenDuration = time.Hour * 24 * 7

func generateTokenPair(data *model.User) (string, string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = data.Uid
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	// Create the refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["uid"] = data.Uid
	refreshClaims["exp"] = time.Now().Add(RefreshTokenDuration).Unix()

	// Sign and get the complete encoded token as a string
	refreshTokenString, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}
