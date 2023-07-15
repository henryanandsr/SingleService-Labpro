package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostBarang(c *gin.Context) {
	var request struct {
		NamaBarang   string `json:"nama"`
		HargaBarang  int    `json:"harga"`
		StokBarang   int    `json:"stok"`
		PerusahaanID string `json:"perusahaan_id"`
		KodeBarang   int    `json:"kode"`
	}
	if request.HargaBarang <= 0 || request.StokBarang < 0 || request.PerusahaanID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data",
			"data":    nil,
		})
		return
	}
	var existingCompany model.Company
	if err := initializers.DB.Where("id = ?", request.PerusahaanID).First(&existingCompany).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Perusahaan not found",
			"data":    nil,
		})
		return
	}
	barang := &model.Barang{
		NamaBarang:   request.NamaBarang,
		HargaBarang:  request.HargaBarang,
		StokBarang:   request.StokBarang,
		IDPerusahaan: existingCompany.Nama,
		KodeBarang:   request.KodeBarang,
	}
	initializers.DB.Create(barang)
}
