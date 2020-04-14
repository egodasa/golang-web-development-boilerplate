package api

import (
    "belajar-ech0-framework/models"
    "github.com/labstack/echo"
    "net/http"
)

// vairabel ini akan dipakai di route
var Mobil = CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
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

func (c CoreApi) Insert(ctx echo.Context) error  {
  data := map[string]interface{}{}
  
  // data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
  // serta data yang ada nilainya
  for _, value := range models.ModelMobil.ColumnList {
    if ctx.FormValue(value.Name) != "" {
      data[value.Name] = string(ctx.FormValue(value.Name));
    }
  }

  // proses insert
  err := models.ModelMobil.Insert(data);
  
  if err == false {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
  }
  
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Update(ctx echo.Context) error  {
  id := ctx.Param("id");
  
  data := map[string]interface{}{}
  
  // data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
  // serta data yang ada nilainya
  for _, value := range models.ModelMobil.ColumnList {
    if ctx.FormValue(value.Name) != "" {
      data[value.Name] = string(ctx.FormValue(value.Name));
    }
  }

  // proses insert
  err := models.ModelMobil.Update(id, data);
  
  if err == false {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
  }
  
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Delete(ctx echo.Context) error  {
  id := ctx.Param("id");
  emptyStruct := struct{}{}
  data := models.ModelMobil.Find(id);
  
  if data == emptyStruct {
    c.HttpStatus = http.StatusNotFound 
    c.Messages = "Data tidak ditemukan"
  } else {
    // proses insert
    err := models.ModelMobil.Delete(id);
    
    if err == false {
      c.HttpStatus = http.StatusInternalServerError
      c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
    }
  }
  
  return ctx.JSON(c.HttpStatus, c)
}


