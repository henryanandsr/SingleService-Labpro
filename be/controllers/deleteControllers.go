package controllers

import (
	repositories "SingleService-Labpro/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	repo := repositories.NewPerusahaanRepository()
	company, err := repo.DeletePerusahaan(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Company and associated barang deleted successfully",
		"data":    company,
	})
}

func DeleteBarang(c *gin.Context) {
	id := c.Param("id")
	repo := repositories.NewBarangRepository()
	barang, err := repo.DeleteBarang(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Barang not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang deleted",
		"data": gin.H{
			"nama":          barang.NamaBarang,
			"harga":         barang.HargaBarang,
			"stok":          barang.StokBarang,
			"kode":          barang.KodeBarang,
			"perusahaan_id": barang.PerusahaanPembuat,
		},
	})
}
