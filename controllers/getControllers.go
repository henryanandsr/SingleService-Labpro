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
func GetPerusahaan(c *gin.Context) {
	id := c.Param("id")

	var perusahaan model.Company
	if err := initializers.DB.Where("ID = ?", id).First(&perusahaan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Company not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Company found",
		"data":    perusahaan,
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

func GetPerusahaans(c *gin.Context) {
	var companies []model.Company
	query := initializers.DB.Model(&model.Company{})

	q := c.Query("q")
	if q != "" {
		query = query.Where("Nama LIKE ? OR KodePajak LIKE ?", "%"+q+"%", "%"+q+"%")
	}

	if err := query.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Companies retrieved successfully",
		"data":    companies,
	})
}
