package routes

import (
	api "golang-web-development/api"
	controllers "golang-web-development/controllers"

	// mdl "golang-web-development/middleware"

	gin "github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var route *gin.Engine = gin.Default()

	// untuk menentukan file statis seperti css, js dll
	// gunakan baris ini jika tidak ingin menggunakan
	// web server pihak ketiga seperti nginx, apache dll
	route.Static("/assets", "./views/assets")

	route.GET("/", controllers.Halaman.Beranda)
	route.GET("/login", controllers.Halaman.Login)
	route.POST("/login", controllers.Halaman.ProsesLogin)
	route.GET("/dashboard", controllers.Halaman.Dashboard)

	var apiRoutes *gin.RouterGroup = route.Group("api")
	{
		apiRoutes.GET("/cek/token", api.Auth.CekToken)
		apiRoutes.POST("/login", api.Auth.CekLogin)

		apiRoutes.GET("/mobilcoba", api.Mobil.CobaMobil)

		// apiRoutes.Use(middleware.ValidasiTokenJwt) // hapus midlleware ini untuk mematikan JWT
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
	}

	return route
}
