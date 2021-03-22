package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	accessSecret = "jdnfksdmfksd"
)

func GenerateToken(username string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expire,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

type tokenVerifyError struct {
	prob string
}

func (err *tokenVerifyError) Error() string {
	return err.prob
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if err := claims.Valid(); err != nil {
			return "", err
		}
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", &tokenVerifyError{"username not found"}
	}
	return "", &tokenVerifyError{"token invalid"}
}
