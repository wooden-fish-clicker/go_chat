package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/wooden-fish-clicker/chat/configs"
)

var jwtSecret = []byte(configs.C.Jwt.Secret)

type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

func (c *Claims) GenerateJWT() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJwtToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			// 檢查token是否過期
			nowTime := time.Now()
			if nowTime.Unix() >= claims.ExpiresAt {
				return nil, errors.New("token已過期")
			}

			// 檢查是否是從我們這裡發出的
			if claims.Issuer != configs.C.App.AppName {
				return nil, errors.New("invalid token issuer")
			}

			return claims, nil
		}
	}

	return nil, err
}
