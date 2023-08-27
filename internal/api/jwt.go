package api

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SimpleSecretKey = "simplesecretkeyblog"

func GenerateJWT(userid int, isAdmin bool, expireDuration time.Duration) (string, error) {
	log.Println(userid, isAdmin)
	dataMap := jwt.MapClaims{
		"userid":    userid,
		"isAdmin":   isAdmin,
		"ExpiredAt": time.Now().Add(expireDuration),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataMap)
	tokenString, err := token.SignedString([]byte(SimpleSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SimpleSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrUnAuthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrUnAuthorized()
	}

	return claims, nil
}
