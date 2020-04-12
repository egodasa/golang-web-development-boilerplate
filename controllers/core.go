package controllers

type CoreController struct {
  HttpStatus int `json:"code"`
  Messages string `json:"message,omitempty"`
  Data interface{} `json:"data,omitempty"`
}
