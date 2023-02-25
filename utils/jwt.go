package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var stSignKey = []byte(viper.GetString("jwt.key"))

type jwtCustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	iJwtCustomClaims := jwtCustomClaims{
		id,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt("jwt.tokenExpire")) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	tokenString, err := token.SignedString(stSignKey)
	return tokenString, err
}

func ParseToken(tokenString string) (jwtCustomClaims, error) {
	iJwtCustomClaims := jwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustomClaims, func(t *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	if err != nil || !token.Valid {
		err = errors.New("invalid token")
	}
	return iJwtCustomClaims, err
}

func IsTokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
