package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	tx := initializers.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": tx.Error.Error(),
		})
		return
	}

	var company model.Company
	if err := tx.Where("ID = ?", id).First(&company).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Company not found",
		})
		return
	}

	var barangs []model.Barang
	if err := tx.Where("PerusahaanPembuat = ?", company.ID).Find(&barangs).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	for _, barang := range barangs {
		if err := tx.Delete(&barang).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Error deleting barang with ID: " + barang.ID,
			})
			return
		}
	}

	if err := tx.Delete(&company).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Transaction failed",
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
	var barang model.Barang
	id := c.Param("id")

	if err := initializers.DB.Where("ID = ?", id).First(&barang).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Barang not found",
			"data":    nil,
		})
		return
	}

	initializers.DB.Delete(&barang)

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
