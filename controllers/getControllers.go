package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBarang(c *gin.Context) {
	id := c.Param("id")

	var barang model.Barang
	if err := initializers.DB.Where("ID = ?", id).First(&barang).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Barang not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Barang found",
		"data":    barang,
	})
}

