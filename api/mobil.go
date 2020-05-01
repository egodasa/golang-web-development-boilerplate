package api

import (
	md "belajar-ech0-framework/models"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

// vairabel ini akan dipakai di route
type ApiMobil struct {
	*CoreApi
}

func (c ApiMobil) Get(ctx *gin.Context) {
	var data, err = md.ModelMobil.GetMobil()
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mengambil data"
	} else {
		c.Data = data
	}
	ctx.JSON(c.HttpStatus, c)
}

var Mobil ApiMobil = ApiMobil{&CoreApi{
	HttpStatus: http.StatusOK,
	Messages:   "Ok",
	ApiModels:  md.ModelMobil,
}}
