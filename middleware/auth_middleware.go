package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(userId int, name string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["isAdmin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("wahyu"))
}
