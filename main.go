package main

import (
  "belajar-ech0-framework/config"
  "belajar-ech0-framework/routes"
  "os"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql" // import your required driver
)

func init() {
  config.LoadConfig();
  DB_CONNECTION := os.Getenv("DB_CONNECTION")
  DB_STRING := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")" + "/" + os.Getenv("DB_DATABASE")
  orm.RegisterDataBase("default", DB_CONNECTION, DB_STRING, 30)
}

func main() {
  config.LoadConfig() // ambil konfigurasi aplikasi
  server := routes.Router() // ambil routes
  // ambil nilai port dari file .env
  port := ":" + os.Getenv("APP_PORT")
  
  server.Start(port)
}
