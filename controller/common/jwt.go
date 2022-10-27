package common

import (
	"errors"
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func JwtAuthentication(ctx *gin.Context) bool {
	tokenHeader := ctx.Request.Header.Get("Authorization") //Grab the token from the header
	if tokenHeader == "" {
		ctx.JSON(common.Failed(common.ErrAuth, nil))
		return false
	}

	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		ctx.JSON(common.Failed(common.ErrAuth, nil))
		return false
	}

	tokenPart := splitted[1] //Grab the token part, what we are truly interested in
	tk := &model.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JwtSecret), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		return false
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		return false
	}

	return true
}

func NewJWTToken(account string) (*model.Token, error) {
	var token = model.Token{
		Account: account,
	}

	token.Issuer = account

	tk := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), token)
	tokenString, _ := tk.SignedString([]byte(config.Cfg.JwtSecret))
	token.Token = tokenString

	return &token, nil
}

func VerifyPassword(param, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(param), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("密码错误")
	}
	return nil
}

func GeneratePassword(password string) string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd)
}
