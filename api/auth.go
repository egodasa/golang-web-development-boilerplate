package api

import (
	hpl "belajar-ech0-framework/helper"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// vairabel ini akan dipakai di route
var Auth = CoreApi{
	HttpStatus: http.StatusOK,
	Messages:   "Ok",
}

func (c CoreApi) CekLogin(ctx *gin.Context) {
	var claims map[string]interface{}
	claims["username"] = ctx.PostForm("username")

	var token string = hpl.GenerateJwtToken(claims, "Kunci")
	c.Data = map[string]string{
		"token": token,
	}
	ctx.JSON(c.HttpStatus, c)
}
