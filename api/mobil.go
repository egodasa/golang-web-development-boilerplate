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
  data, err := models.ModelMobil.Get();
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan dalam mengambil data"
  } else {
    c.Data = data
  }
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Find(ctx echo.Context) error  {
  id := ctx.Param("id")
  data, err := models.ModelMobil.Find(id);
  
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan dalam mencari data"
  } else {
    c.Data = data
  }
  
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
  id, err := models.ModelMobil.Insert(data);
  
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan saat menambahkan data"
  } else {
    data := map[string]string{}
    data["id"] = id
    c.Data = data
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
  update := models.ModelMobil.Update(id, data);
  
  if update == false {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan saat mengupdate data"
  }
  
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Delete(ctx echo.Context) error  {
  id := ctx.Param("id");
  emptyStruct := struct{}{}
  data, err := models.ModelMobil.Find(id);
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan dalam menghapus data"
  } else {
      if data == emptyStruct {
        c.HttpStatus = http.StatusNotFound 
        c.Messages = "Data tidak ditemukan"
      } else {
        // proses insert
        delete := models.ModelMobil.Delete(id);
        
        if delete == false {
          c.HttpStatus = http.StatusInternalServerError
          c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
        }
      }
  }
  
  
  return ctx.JSON(c.HttpStatus, c)
}


