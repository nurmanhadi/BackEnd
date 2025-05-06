package service

import (
	"liva/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtGenerateToken(id string, key []byte) (string, error) {
	claims := model.JwtCustomClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 24) * 7)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", err
	}
	return token, nil
}
func JwtVerifyToken(tokenString string, key []byte) (*model.JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*model.JwtCustomClaims)
	claimType := &model.JwtCustomClaims{
		Id: claims.Id,
	}
	return claimType, nil
}
