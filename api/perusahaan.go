package api

import (
    md "belajar-ech0-framework/models"
    "net/http"
)

// vairabel ini akan dipakai di route
var Perusahaan = CoreApi{
  HttpStatus: http.StatusOK,
  Messages: "Ok",
  ApiModels: md.ModelPerusahaan,
}



