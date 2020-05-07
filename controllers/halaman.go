package controllers

import (
	"bytes"
	views "golang-web-development/views/dist"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// buat objek halaman controller
var Halaman = CoreController{
	HttpStatus: http.StatusOK,
}

func (c CoreController) Beranda(ctx *gin.Context) {
	buffer := new(bytes.Buffer)
	views.Beranda(buffer)

	c.ViewHTML(ctx, c.HttpStatus, buffer.Bytes())
}

func (c CoreController) Login(ctx *gin.Context) {
	buffer := new(bytes.Buffer)
	views.Login(buffer)
	c.ViewHTML(ctx, c.HttpStatus, buffer.Bytes())
}

func (c CoreController) ProsesLogin(ctx *gin.Context) {
	var username string = ctx.PostForm("username")
	var password string = ctx.PostForm("password")
	var data string

	if username != "madam" && password != "12345" {
		data = "Username atau password salah <br><a href='login'>Kembali</a>"
		c.HttpStatus = http.StatusNonAuthoritativeInfo
		ctx.HTML(c.HttpStatus, data, nil)
	} else {
		data = "/dashboard"
		c.HttpStatus = http.StatusFound
		ctx.Redirect(c.HttpStatus, data)
	}
}

func (c CoreController) Dashboard(ctx *gin.Context) {
	var data string = "<h1>Selamat datang didashboard</h1>"

	c.HttpStatus = http.StatusOK

	ctx.HTML(c.HttpStatus, data, nil)
}
