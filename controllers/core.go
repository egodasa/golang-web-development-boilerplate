package controllers

import (
  "net/http"
  "github.com/labstack/echo"
)

type CoreController struct {
  httpStatus int
  error string
  //~ data interface{}
}

func (c CoreController) HalamanError(ctx echo.Context) error {
  return ctx.HTML(c.httpStatus, c.error)
}
func (c CoreController) HalamanUnderConstruction(ctx echo.Context) error {
  data := "Halaman belum siap!"
  c.httpStatus = http.StatusOK
  return ctx.HTML(c.httpStatus, data)
}
