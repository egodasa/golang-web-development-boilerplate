package api

import (
    md "belajar-ech0-framework/models"
    "net/http"
)

// vairabel ini akan dipakai di route
type ApiMobil struct {
  *CoreApi
}

var Mobil ApiMobil = ApiMobil{&CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
  ApiModels: md.ModelMobil,
}}



