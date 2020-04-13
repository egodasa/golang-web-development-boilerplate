package api

import (
    "belajar-ech0-framework/models"
    "github.com/labstack/echo"
    "net/http"
)

// vairabel ini akan dipakai di route
var Mobil = CoreApi{
  HttpStatus: http.StatusOK,
}

func (c CoreApi) Get(ctx echo.Context) error  {
  c.Data = models.ModelMobil.Get();
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Find(ctx echo.Context) error  {
  id := ctx.Param("id")
  c.Data = models.ModelMobil.Find(id);
  return ctx.JSON(c.HttpStatus, c)
}

