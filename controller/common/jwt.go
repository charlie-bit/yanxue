package common

import (
	"context"
	"errors"
	"github.com/charlie-bit/yanxue/config"
	"github.com/charlie-bit/yanxue/db"
	"github.com/charlie-bit/yanxue/model"
	"github.com/charlie-bit/yanxue/model/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func JwtAuthentication(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization") //Grab the token from the header
	if tokenHeader == "" {
		ctx.AbortWithStatusJSON(common.Failed(common.ErrAuth, nil))
		return
	}

	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		ctx.AbortWithStatusJSON(common.Failed(common.ErrAuth, nil))
		return
	}

	tokenPart := splitted[1] //Grab the token part, what we are truly interested in
	tk := &model.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JwtSecret), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		ctx.AbortWithStatusJSON(common.Failed(common.ErrAuth, err))
		return
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		ctx.AbortWithStatusJSON(common.Failed(common.ErrAuth, nil))
		return
	}

	account := ctx.Request.Header.Get("account")

	if tokenPart != db.RedisClient.Get(context.Background(), account).Val() {
		ctx.AbortWithStatusJSON(common.Failed(common.ErrAuth, nil))
		return
	}

	ctx.Next()
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
