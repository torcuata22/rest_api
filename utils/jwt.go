package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superdupersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), //token valid for 2 hours, Unix format
	})
	return token.SignedString([]byte(secretKey))
}
