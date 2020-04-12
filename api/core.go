package api

type CoreApi struct {
  HttpStatus int `json:"code"` // field untuk kode http
  Messages string `json:"message,omitempty"` // field untuk pesan
  Data interface{} `json:"data,omitempty"` // field yang menamping data yang akan diberikan ke klien
}

