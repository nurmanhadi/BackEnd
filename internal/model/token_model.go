package model

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Id          string `json:"id"`
	RerefenceId string `json:"reference_id"`
	Role        string `json:"role"`
	jwt.RegisteredClaims
}
