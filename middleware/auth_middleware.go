package middleware

import (
	"github.com/golang-jwt/jwt"
	"mini_project/constant"
	"time"
)

func CreateTokenForUser(userId int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.USER_SECRET_JWT))
}

func CreateTokeForAdmin(adminId int, touristAttractionID int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["adminId"] = adminId
	claims["touristAttractionID"] = touristAttractionID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.ADMIN_SECRET_JWT))
}
