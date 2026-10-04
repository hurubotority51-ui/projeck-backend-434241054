package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("vehicle-rental-secret-key")

func GenerateAccessToken(userID int, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("gagal membuat access token: %w", err)
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("algoritma token tidak valid")
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("token tidak valid: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims token tidak valid")
	}

	return claims, nil
}