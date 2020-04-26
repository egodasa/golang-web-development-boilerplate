package api

import (
    md "belajar-ech0-framework/models"
    "net/http"
    
    gin "github.com/gin-gonic/gin"
)

// vairabel ini akan dipakai di route
type ApiJenisMobil struct {
  *CoreApi
}

func (c ApiJenisMobil) Cari(ctx *gin.Context) {
  modelMobil := md.ModelJenisMobil
	var id string = ctx.Param("cari")
	var data, err = modelMobil.Cari(id)

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mencari data"
	} else {
		c.Data = data
	}

	ctx.JSON(c.HttpStatus, c)
}

var JenisMobil ApiJenisMobil = ApiJenisMobil{&CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
  ApiModels: md.ModelJenisMobil,
}}



