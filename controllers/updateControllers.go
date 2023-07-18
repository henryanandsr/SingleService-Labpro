package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateBarang(c *gin.Context) {
	id := c.Param("id")
	var barang model.Barang

	if err := initializers.DB.Where("ID = ?", id).First(&barang).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Barang not found",
		})
		return
	}

	var updateData struct {
		NamaBarang        string `json:"namaBarang"`
		HargaBarang       int    `json:"hargaBarang"`
		StokBarang        int    `json:"stokBarang"`
		PerusahaanPembuat string `json:"perusahaanPembuat"`
		KodeBarang        string `json:"kodeBarang"`
	}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	barang.NamaBarang = updateData.NamaBarang
	barang.HargaBarang = updateData.HargaBarang
	barang.StokBarang = updateData.StokBarang
	barang.PerusahaanPembuat = updateData.PerusahaanPembuat
	barang.KodeBarang = updateData.KodeBarang

	if err := initializers.DB.Save(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang updated successfully",
		"data":    barang,
	})
}

func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company model.Company

	if err := initializers.DB.Where("ID = ?", id).First(&company).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Company not found",
		})
		return
	}

	var updateData struct {
		Nama      string `json:"nama"`
		Alamat    string `json:"alamat"`
		NoTelepon string `json:"noTelepon"`
		KodePajak string `json:"kodePajak"`
	}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	company.Nama = updateData.Nama
	company.Alamat = updateData.Alamat
	company.NoTelepon = updateData.NoTelepon
	company.KodePajak = updateData.KodePajak

	if err := initializers.DB.Save(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Company updated successfully",
		"data":    company,
	})
}
