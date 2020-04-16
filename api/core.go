package api

import (
  md "belajar-ech0-framework/models"
  "github.com/labstack/echo"
  "net/http"
)

type CoreApi struct {
  HttpStatus int `json:"code"` // field untuk kode http
  Messages string `json:"message,omitempty"` // field untuk pesan
  Data interface{} `json:"data,omitempty"` // field yang menamping data yang akan diberikan ke klien
  ApiModels *md.Models `json:"-"`
}

func (c CoreApi) Get(ctx echo.Context) error  {
  data, err := c.ApiModels.Get();
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
  data, err := c.ApiModels.Find(id);
  
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan dalam mencari data"
  } else {
    c.Data = data
  }
  
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Insert(ctx echo.Context) error  {
  data := make(map[string]string)
  
  // data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
  // serta data yang ada nilainya
  for _, value := range c.ApiModels.ColumnList {
    if ctx.FormValue(value.Name) != "" {
      data[value.Name] = string(ctx.FormValue(value.Name));
    }
  }

  // proses insert
  id, err := c.ApiModels.Insert(data);
  
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
  
  data := make(map[string]string)
  
  // data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
  // serta data yang ada nilainya
  for _, value := range c.ApiModels.ColumnList {
    if ctx.FormValue(value.Name) != "" {
      data[value.Name] = ctx.FormValue(value.Name);
    }
  }

  // proses insert
  update := c.ApiModels.Update(id, data);
  
  if update == false {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan saat mengupdate data"
  }
  
  return ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Delete(ctx echo.Context) error  {
  id := ctx.Param("id");
  data, err := c.ApiModels.Find(id);
  if err == true {
    c.HttpStatus = http.StatusInternalServerError
    c.Messages = "Terjadi kesalahan dalam menghapus data"
  } else {
      if data == nil {
        c.HttpStatus = http.StatusNotFound 
        c.Messages = "Data tidak ditemukan"
      } else {
        // proses insert
        delete := c.ApiModels.Delete(id);
        
        if delete == false {
          c.HttpStatus = http.StatusInternalServerError
          c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
        }
      }
  }
  
  
  return ctx.JSON(c.HttpStatus, c)
}

