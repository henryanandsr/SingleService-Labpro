package main

import (
	"SingleService-Labpro/controllers"
	"SingleService-Labpro/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables(".env")
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/barang", controllers.PostBarang)
	r.POST("/perusahaan", controllers.PostCompany)
	r.DELETE("/perusahaan/:id", controllers.DeleteCompany)
	r.DELETE("/barang/:id", controllers.DeleteBarang)
	r.GET("/barang/:id", controllers.GetBarang)
	r.GET("/perusahaan/:id", controllers.GetPerusahaan)
	r.GET("/perusahaan", controllers.GetPerusahaans)
	r.GET("/barang", controllers.GetBarangs)
	r.Run(":8080")
}
