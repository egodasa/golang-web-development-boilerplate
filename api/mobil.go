package api

import (
    "belajar-ech0-framework/models"
    "github.com/labstack/echo"
    "net/http"
    "strconv"
)

// vairabel ini akan dipakai di route
var Mobil = CoreApi{
  HttpStatus: http.StatusOK,
}

var modelMobil = models.ModelMobil{}

func (c CoreApi) Get(ctx echo.Context) error  {
  c.Data = modelMobil.Get();
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Find(ctx echo.Context) error  {
  id, err := strconv.Atoi(ctx.Param("id"))
  if err != nil {
    c.HttpStatus = http.StatusNotFound
    c.Messages = "Halaman tidak ditemukan"
    c.Data = nil
  } else {
    c.Data = modelMobil.Find(id);
  }
  return ctx.JSON(c.HttpStatus, c)
}

