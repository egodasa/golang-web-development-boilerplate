package routes

import (
  "belajar-ech0-framework/controllers"
  "belajar-ech0-framework/api"
  "github.com/labstack/echo"
)

func Router() *echo.Echo {
    r := echo.New()

    r.GET("/", controllers.Halaman.Beranda)
    r.GET("/login", controllers.Halaman.Login)
    r.POST("/login", controllers.Halaman.ProsesLogin)
    r.GET("/dashboard", controllers.Halaman.Dashboard)
    
    r.GET("/api/mobil", api.Mobil.Get)
    r.GET("/api/mobil/:id", api.Mobil.Find)
    r.POST("/api/mobil", api.Mobil.Insert)
    r.PUT("/api/mobil/:id", api.Mobil.Update)
    r.DELETE("/api/mobil/:id", api.Mobil.Delete)
    return r
}
