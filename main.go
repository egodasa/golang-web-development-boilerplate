package main

import (
  "belajar-ech0-framework/controllers"
  "belajar-ech0-framework/api"
  "github.com/labstack/echo"
  "github.com/joho/godotenv"
  "os"
  "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
)

func Router() *echo.Echo {
    r := echo.New()

    r.GET("/", controllers.Halaman.Beranda)
    r.GET("/login", controllers.Halaman.Login)
    r.POST("/login", controllers.Halaman.ProsesLogin)
    r.GET("/dashboard", controllers.Halaman.Dashboard)
    
    r.GET("/api/mobil", api.Mobil.Get)
    r.GET("/api/mobil/:id", api.Mobil.Find)
    return r
}

func main() {
  flag := kingpin.New("Aplikasi CRUD", "Aplikasi CRUD dengan echo framework dan golang")
  
  // menagkap flag --env=".env"
  flagEnv := flag.Flag("env", "Lokasi file .env (environment) atau file custom .env").Default(".env").String() // required = flag wajib ada, string = isi flag harus string pakai kutip
  kingpin.MustParse(flag.Parse(os.Args[1:]))
  
  // ambil file env yang berisi konfigurasi
  err := godotenv.Load(*flagEnv)
  if err != nil {
    // jika file .env atau semacamnya tidak ditemukan, stop jalan aplikasi
    fmt.Println("Error pada aplikasi!")
    fmt.Println("File environment " + *flagEnv + " tidak ditemukan!")
    return
  }
  
  // ambil nilai port dari file .env
  port := ":" + os.Getenv("APP_PORT")
  
  r := Router(); // method ini berasal dari file routes.go
  r.Start(port)
}
