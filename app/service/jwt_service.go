package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID int, username string, role string) (string, error) {
	secret, err := loadJWTSecret()
	if err != nil {
		return "", err
	}

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

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("gagal membuat access token: %w", err)
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (jwt.MapClaims, error) {
	secret, err := loadJWTSecret()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("algoritma token tidak valid")
			}

			return secret, nil
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

func loadJWTSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("JWT_SECRET belum diatur")
	}

	return []byte(secret), nil
}