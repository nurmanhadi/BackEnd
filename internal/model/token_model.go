package model

import "github.com/golang-jwt/jwt/v5"

type TokenResponse struct {
	Token string `json:"token"`
}
type JwtCustomClaims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}
