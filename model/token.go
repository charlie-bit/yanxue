package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	jwt.StandardClaims
	Account string
	Token   string
}
