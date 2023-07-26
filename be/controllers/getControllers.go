package controllers

import (
	repositories "SingleService-Labpro/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBarang(c *gin.Context) {
	id := c.Param("id")
	repo := repositories.NewBarangRepository()
	barang, err := repo.GetBarang(id)
	if err != nil {
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
	q := c.Query("q")
	perusahaan := c.Query("perusahaan")
	repo := repositories.NewBarangRepository()
	barangs, err := repo.GetAllBarangs(q, perusahaan)
	if err != nil {
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
func GetPerusahaan(c *gin.Context) {
	id := c.Param("id")
	repo := repositories.NewPerusahaanRepository()
	perusahaan, err := repo.GetPerusahaan(id)
	if err != nil {
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

func GetPerusahaans(c *gin.Context) {
	q := c.Query("q")
	repo := repositories.NewPerusahaanRepository()
	companies, err := repo.GetAllPerusahaans(q)
	if err != nil {
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
