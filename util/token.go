package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mini_project/constant"
)

func ParsingToken(token string) (*jwt.Token, error) {
	result, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constant.USER_SECRET_JWT), nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
