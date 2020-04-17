package api

import (
    "net/http"
    hpl "belajar-ech0-framework/helper"
    
    "github.com/labstack/echo"
)

// vairabel ini akan dipakai di route
var Auth = CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
}

func (c CoreApi) CekLogin(ctx echo.Context) error  {
  claims := map[string]interface{}{
    "username": ctx.FormValue("username"),
  }
  
  token := hpl.GenerateJwtToken(claims, "Kunci");
  c.Data = map[string]string{
    "token": token,
  }
  return ctx.JSON(c.HttpStatus, c)
}

