package controllers

import (
	gin "github.com/gin-gonic/gin"
)

type CoreController struct {
	HttpStatus int         `json:"code"`
	Messages   string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (c *CoreController) ViewHTML(ctx *gin.Context, status int, html []byte) {
	ctx.Data(status, "text/html; charset=utf-8", html)
}
