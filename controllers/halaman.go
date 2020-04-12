package controllers

import (
    "github.com/labstack/echo"
    "net/http"
)

// buat objek halaman controller
var Halaman = CoreController{
  httpStatus: http.StatusOK,
}

func (c CoreController) Beranda(ctx echo.Context) error  {
  data := "Ini halaman beranda"
  return ctx.String(c.httpStatus, data)
}

func (c CoreController) Login(ctx echo.Context) error  {
  data := `
            <h1>Ini halaman login</h1>
            <form method="POST" action="">
              <label>Username</label><br>
              <input type="text" name="username" /><br>
              
              <label>Password</label><br>
              <input type="password" name="password" /><br>
              
              <button type="submit">Login</button>
            </form>
          `
  return ctx.HTML(c.httpStatus, data)
}

func (c CoreController) ProsesLogin(ctx echo.Context) error  {
  username :=  ctx.FormValue("username");
  password :=  ctx.FormValue("password");

  if username != "madam" && password != "12345" {
    data := "Username atau password salah"
    data += "<br><a href='login'>Kembali</a>"
    c.httpStatus = http.StatusNonAuthoritativeInfo
    return ctx.HTML(c.httpStatus, data)
  } else {
    data := "/dashboard"
    c.httpStatus = http.StatusFound
    return ctx.Redirect(c.httpStatus, data)
  }
}

func (c CoreController) Dashboard(ctx echo.Context) error  {
  c.httpStatus = http.StatusOK
  data := "<h1>Selamat datang didashboard</h1>"
  
  return ctx.HTML(c.httpStatus, data)
}

