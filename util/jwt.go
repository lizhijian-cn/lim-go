package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secret = "lim"
)

func GenerateToken(username string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expire,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if err := claims.Valid(); err != nil {
			return "", err
		}
		if username, ok := claims["username"].(string); !ok {
			err = errors.New("failed to parse username in token")
		} else {
			return username, nil
		}
	}
	return "", errors.New("invalid token")
}
