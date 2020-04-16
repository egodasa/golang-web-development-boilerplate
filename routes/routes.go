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
    
    r.GET("/api/perusahaan", api.Perusahaan.Get)
    r.GET("/api/perusahaan/:id", api.Perusahaan.Find)
    r.POST("/api/perusahaan", api.Perusahaan.Insert)
    r.PUT("/api/perusahaan/:id", api.Perusahaan.Update)
    r.DELETE("/api/perusahaan/:id", api.Perusahaan.Delete)
    
    r.GET("/api/jenis-mobil", api.JenisMobil.Get)
    r.GET("/api/jenis-mobil/:id", api.JenisMobil.Find)
    r.POST("/api/jenis-mobil", api.JenisMobil.Insert)
    r.PUT("/api/jenis-mobil/:id", api.JenisMobil.Update)
    r.DELETE("/api/jenis-mobil/:id", api.JenisMobil.Delete)
    return r
}
