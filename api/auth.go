package api

import (
	hpl "golang-web-development/helper"
	http "net/http"
	os "os"
	strings "strings"
	time "time"

	gin "github.com/gin-gonic/gin"
)

// vairabel ini akan dipakai di route
type ApiAuth struct {
	*CoreApi
}

var Auth ApiAuth = ApiAuth{&CoreApi{
	HttpStatus: http.StatusOK,
	Messages:   "Ok",
}}

// CekLogin
// Controller untuk generate token JWT dari data login
func (c ApiAuth) CekLogin(ctx *gin.Context) {
	claims := make(map[string]interface{})
	claims["iss"] = time.Now().Unix()
	claims["username"] = ctx.PostForm("username")

	var token string = hpl.GenerateJwtToken(claims, os.Getenv("JWT_KEY"))
	c.Data = map[string]string{
		"token": token,
	}
	ctx.JSON(c.HttpStatus, c)
}

// CekToken
// Controller untuk cek token berdasarkan token yang dikirim lewat
// header Authorization
func (c CoreApi) CekToken(ctx *gin.Context) {
	var authorizationHeader = ctx.Request.Header.Get("Authorization")
	if authorizationHeader == "" {
		c.HttpStatus = http.StatusForbidden
		c.Messages = "Akses ditolak"
	} else {
		var token = strings.Split(authorizationHeader, " ")
		if len(token) != 2 {
			c.HttpStatus = http.StatusInternalServerError
			c.Messages = "Token tidak valid"
		} else {
			if token[0] != "Bearer" {
				c.HttpStatus = http.StatusInternalServerError
				c.Messages = "Token tidak valid"
			} else {
				if hpl.ValidateJwtToken(token[1], os.Getenv("JWT_KEY")) == true {
					c.Messages = "Token valid"
				} else {
					c.HttpStatus = http.StatusInternalServerError
					c.Messages = "Token tidak valid"
				}
			}
		}

	}
	ctx.JSON(c.HttpStatus, c)
}
