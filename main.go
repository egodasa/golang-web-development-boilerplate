package main

import (
    "belajar-ech0-framework/controllers"
    "github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
    r := echo.New()

    r.GET("/", controllers.Halaman.Beranda)
    r.GET("/login", controllers.Halaman.Login)
    r.POST("/login", controllers.Halaman.ProsesLogin)
    r.GET("/dashboard", controllers.Halaman.Dashboard)

    r.Start(":9000")
}
