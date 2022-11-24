package auth

import (
	"dormitory-system/src/model"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const TokenDuration = time.Minute * 5
const RefreshTokenDuration = time.Hour * 24 * 7

func generateTokenPair(data *model.User) (string, string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = data.Uid
	claims["exp"] = time.Now().Add(TokenDuration).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
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

func ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if token.Valid && err == nil {
		claims := token.Claims.(jwt.MapClaims)
		uid := int(claims["uid"].(float64))
		return uid, nil
	} else {
		return 0, err
	}
}
