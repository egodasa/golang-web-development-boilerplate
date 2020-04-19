package main

import (
	config "belajar-ech0-framework/config"
	routes "belajar-ech0-framework/routes"
	os "os"
  
  gin "github.com/gin-gonic/gin"
	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

func init() {
	config.LoadConfig()
	var dbConnection string = os.Getenv("DB_CONNECTION")
	var dbString string = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")" + "/" + os.Getenv("DB_DATABASE")
	orm.RegisterDataBase("default", dbConnection, dbString, 30)
}

func main() {
  config.LoadConfig()// ambil konfigurasi aplikasi
  var server *gin.Engine = routes.Router() // ambil routes
  var port string = ":" + os.Getenv("APP_PORT")
	server.Run(port)
}
