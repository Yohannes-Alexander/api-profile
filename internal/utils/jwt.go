package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func GenerateTokens(userID int64) (string, string, error) {
	accessToken, err := generateToken(userID, 15*time.Minute, false)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateToken(userID, 7*24*time.Hour, true)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func generateToken(userID int64, duration time.Duration, isRefresh bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
		"type":    "access",
	}
	if isRefresh {
		claims["type"] = "refresh"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenStr string, expectRefresh bool) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["type"] != "refresh" && expectRefresh {
		return 0, errors.New("not a refresh token")
	}
	userID := int64(claims["user_id"].(float64))
	return userID, nil
}
