package controllers

import (
    gin "github.com/gin-gonic/gin"
    http "net/http"
)

// buat objek halaman controller
var Halaman = CoreController{
  HttpStatus: http.StatusOK,
}

func (c CoreController) Beranda(ctx *gin.Context) {
  var data string = "Ini halaman beranda"
  ctx.String(c.HttpStatus, data)
}

func (c CoreController) Login(ctx *gin.Context) {
  var data string = `
            <h1>Ini halaman login</h1>
            <form method="POST" action="">
              <label>Username</label><br>
              <input type="text" name="username" /><br>
              
              <label>Password</label><br>
              <input type="password" name="password" /><br>
              
              <button type="submit">Login</button>
            </form>
          `
  ctx.HTML(c.HttpStatus, data)
}

func (c CoreController) ProsesLogin(ctx *gin.Context) {
  var username string =  ctx.PostForm("username");
  var password string =  ctx.PostForm("password");
  var data string
  
  if username != "madam" && password != "12345" {
    data = "Username atau password salah <br><a href='login'>Kembali</a>"
    c.HttpStatus = http.StatusNonAuthoritativeInfo
    ctx.HTML(c.HttpStatus, data)
  } else {
    data = "/dashboard"
    c.HttpStatus = http.StatusFound
    ctx.Redirect(c.HttpStatus, data)
  }
}

func (c CoreController) Dashboard(ctx *gin.Context) {
  var data string = "<h1>Selamat datang didashboard</h1>"
  
  c.HttpStatus = http.StatusOK
  
  ctx.HTML(c.HttpStatus, data)
}


