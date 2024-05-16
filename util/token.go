package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mini_project/constant"
	"strings"
)

func parseToken(token string) (*jwt.Token, error) {

	tokenString := strings.Split(token, " ")[1]

	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.USER_SECRET_JWT), nil
	})
	if err != nil {
		return nil, constant.ErrInvalidToken
	}
	return result, nil
}

func GetUserIdFromToken(token string) (int, error) {
	result, err := parseToken(token)
	if err != nil {
		return 0, err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := int(claims["userId"].(float64))

		return userID, err
	}
	return 0, constant.ErrInvalidToken
}

func GetTouristAttractionIdFromToken(token string) (int, error) {
	result, err := parseToken(token)
	if err != nil {
		return 0, err
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := int(claims["touristAttractionID"].(float64))

		return userID, err
	}
	return 0, constant.ErrInvalidToken
}
