package main

import (
  "belajar-ech0-framework/config"
)

func main() {
  config.LoadConfig() // ambil konfigurasi aplikasi
  r := config.Router() // ambil routes
  // ambil nilai port dari file .env
  port := ":" + os.Getenv("APP_PORT")
  
  r := Router(); // method ini berasal dari file routes.go
  r.Start(port)
}
