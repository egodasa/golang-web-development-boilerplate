package main

import (
  "belajar-ech0-framework/controllers"
  "github.com/labstack/echo"
  "github.com/joho/godotenv"
  "os"
)

func Router() *echo.Echo {
    r := echo.New()

    r.GET("/", controllers.Halaman.Beranda)
    r.GET("/login", controllers.Halaman.Login)
    r.POST("/login", controllers.Halaman.ProsesLogin)
    r.GET("/dashboard", controllers.Halaman.Dashboard)
    return r
}

func main() {
  // ambil file env yang berisi konfigurasi
  err := godotenv.Load()
  if err != nil {
    panic("File .env not found")
  }
  
  // ambil nilai port dari file .env
  port := ":" + os.Getenv("APP_PORT")
  
  r := Router(); // method ini berasal dari file routes.go
  r.Start(port)
}
