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
	r.POST("/company", controllers.PostCompany)
	r.DELETE("/company/:id", controllers.DeleteCompany)
	r.Run(":8080")
}
