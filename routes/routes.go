package routes

import (
	"belajar-ech0-framework/api"
	"belajar-ech0-framework/controllers"
	mdl "belajar-ech0-framework/middleware"

	"github.com/labstack/echo"
)

func Router() *echo.Echo {
	r := echo.New()

	r.GET("/", controllers.Halaman.Beranda)
	r.GET("/login", controllers.Halaman.Login)
	r.POST("/login", controllers.Halaman.ProsesLogin)
	r.GET("/dashboard", controllers.Halaman.Dashboard)

	apiRoutes := r.Group("api")
  
	apiRoutes.POST("/login", api.Auth.CekLogin)
  
  apiRoutes.Use(mdl.CekLogin);
  
	apiRoutes.GET("/mobil", api.Mobil.Get)
	apiRoutes.GET("/mobil/:id", api.Mobil.Find)
	apiRoutes.POST("/mobil", api.Mobil.Insert)
	apiRoutes.PUT("/mobil/:id", api.Mobil.Update)
	apiRoutes.DELETE("/mobil/:id", api.Mobil.Delete)
  
	apiRoutes.GET("/perusahaan", api.Perusahaan.Get)
	apiRoutes.GET("/perusahaan/:id", api.Perusahaan.Find)
	apiRoutes.POST("/perusahaan", api.Perusahaan.Insert)
	apiRoutes.PUT("/perusahaan/:id", api.Perusahaan.Update)
	apiRoutes.DELETE("/perusahaan/:id", api.Perusahaan.Delete)
  
	apiRoutes.GET("/jenis-mobil", api.JenisMobil.Get)
	apiRoutes.GET("/jenis-mobil/:id", api.JenisMobil.Find)
	apiRoutes.POST("/jenis-mobil", api.JenisMobil.Insert)
	apiRoutes.PUT("/jenis-mobil/:id", api.JenisMobil.Update)
	apiRoutes.DELETE("/jenis-mobil/:id", api.JenisMobil.Delete)
	return r
}
