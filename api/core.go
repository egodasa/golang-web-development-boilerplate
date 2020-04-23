package api

import (
	md "belajar-ech0-framework/models"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

type CoreApi struct {
	HttpStatus int         `json:"code"`              // field untuk kode http
	Messages   string      `json:"message,omitempty"` // field untuk pesan
	Data       interface{} `json:"data,omitempty"`    // field yang menamping data yang akan diberikan ke klien
	ApiModels  *md.Models  `json:"-"`
}

func (c CoreApi) Get(ctx *gin.Context) {
	var data, err = c.ApiModels.Get()
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mengambil data"
	} else {
		c.Data = data
	}
	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Find(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var data, err = c.ApiModels.Find(id)

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam mencari data"
	} else {
		c.Data = data
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Insert(ctx *gin.Context) {
	data :=  make(map[string]string)

	// data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
	// serta data yang ada nilainya
	for _, value := range c.ApiModels.ColumnList {
		if ctx.PostForm(value.Name) != "" {
			data[value.Name] = string(ctx.PostForm(value.Name))
		}
	}

	// proses insert
	var id, err = c.ApiModels.Insert(data)

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan saat menambahkan data"
	} else {
		data := make(map[string]string)
		data["id"] = id
		c.Data = data
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Update(ctx *gin.Context) {
	var id string = ctx.Param("id")

	data :=  make(map[string]string)

	// data yang dimasukkan hanyalah data yang sudah ditentukan di ColumnList
	// serta data yang ada nilainya
	for _, value := range c.ApiModels.ColumnList {
		if ctx.PostForm(value.Name) != "" {
			data[value.Name] = ctx.PostForm(value.Name)
		}
	}

	// proses insert
	err := c.ApiModels.Update(id, data)

	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan saat mengupdate data"
	}

	ctx.JSON(c.HttpStatus, c)
}

func (c CoreApi) Delete(ctx *gin.Context) {
	var id string = ctx.Param("id")
	var data, err = c.ApiModels.Find(id)
	if err == true {
		c.HttpStatus = http.StatusInternalServerError
		c.Messages = "Terjadi kesalahan dalam menghapus data"
	} else {
		if data == nil {
			c.HttpStatus = http.StatusNotFound
			c.Messages = "Data tidak ditemukan"
		} else {
			// proses insert
			err = c.ApiModels.Delete(id)

			if err == true {
				c.HttpStatus = http.StatusInternalServerError
				c.Messages = "Terjadi kesalahan. Tidak dapat menyimpan data!"
			}
		}
	}

	ctx.JSON(c.HttpStatus, c)
}
