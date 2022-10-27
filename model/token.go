package model

import "github.com/golang-jwt/jwt"

type Token struct {
	jwt.StandardClaims
	Account string
	Token   string
}
