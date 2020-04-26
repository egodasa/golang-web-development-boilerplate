package api

import (
    md "belajar-ech0-framework/models"
    "net/http"
)

// vairabel ini akan dipakai di route
type ApiPerusahaan struct {
  *CoreApi
}

var Perusahaan ApiPerusahaan = ApiPerusahaan{&CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
  ApiModels: md.ModelPerusahaan,
}}



