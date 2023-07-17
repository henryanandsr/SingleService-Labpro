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

func GetBarangs(c *gin.Context) {
	var barangs []model.Barang
	query := initializers.DB.Model(&model.Barang{})

	q := c.Query("q")
	if q != "" {
		query = query.Where("NamaBarang LIKE ? OR KodeBarang LIKE ?", "%"+q+"%", "%"+q+"%")
	}

	perusahaan := c.Query("perusahaan")
	if perusahaan != "" {
		query = query.Where("PerusahaanPembuat = ?", perusahaan)
	}

	if err := query.Find(&barangs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barangs retrieved successfully",
		"data":    barangs,
	})
}
