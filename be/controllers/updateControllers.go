package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateStokBarang(c *gin.Context) {
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
		StokBarang int `json:"stok"` // only allow updating stock
	}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	barang.StokBarang = updateData.StokBarang

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

func UpdateBarang(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := authHeader
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if err != nil || !tkn.Valid || claims.Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
		return
	}
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
		NamaBarang        string `json:"nama"`
		HargaBarang       int    `json:"harga"`
		StokBarang        int    `json:"stok"`
		PerusahaanPembuat string `json:"perusahaan_id"`
		KodeBarang        string `json:"kode"`
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
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := authHeader
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})
	if err != nil || !tkn.Valid || claims.Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
		return
	}
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
		NoTelepon string `json:"no_telp"`
		KodePajak string `json:"kode"`
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
