package api

import (
	md "belajar-ech0-framework/models"
	"math"
	"net/http"

	"strconv"

	gin "github.com/gin-gonic/gin"
)

// vairabel ini akan dipakai di route
type ApiMobil struct {
	*CoreApi
}

func (c ApiMobil) Get(ctx *gin.Context) {

	// set limit dan offset untuk paginasi data
	page, errPage := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if errPage != nil {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Nilai query PAGE tidak valid!"
		ctx.AbortWithStatusJSON(c.HttpStatus, c)
	}
	perPage, errPerPage := strconv.Atoi(ctx.DefaultQuery("perpage", "10"))
	if errPerPage != nil {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Nilai query PERPAGE tidak valid!"
		ctx.AbortWithStatusJSON(c.HttpStatus, c)
	}

	// ambil total data secara keseluruhan tanpa limit
	totalData, _ := md.ModelMobil.Count()

	// hitung total halaman
	totalPage := math.Ceil(float64(totalData) / float64(perPage))

	// limit untuk query sql
	var limit int = 10

	// jika perpage diset ke 0, maka semua data akan diambil
	if perPage == 0 {
		page = 0
	} else {
		if page > 1 {
			limit = (page - 1) * perPage
		}
	}

	var data, err = md.ModelMobil.GetMobil(limit, perPage)
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mengambil data"
	} else {
		c.Data = map[string]interface{}{
			"data":         data,
			"current_page": page,
			"per_page":     perPage,
			"total_page":   int(totalPage),
			"total_data":   totalData,
		}
	}
	ctx.JSON(c.HttpStatus, c)
}

var Mobil ApiMobil = ApiMobil{&CoreApi{
	HttpStatus: http.StatusOK,
	Messages:   "Ok",
	ApiModels:  md.ModelMobil,
}}
