package controllers

import (
	repositories "SingleService-Labpro/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateStokBarang(c *gin.Context) {
	id := c.Param("id")
	var updateData struct {
		StokBarang int `json:"stok"` // only update stok
	}

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	repo := repositories.NewBarangRepository()
	barang, err := repo.UpdateStokBarang(id, updateData.StokBarang)
	if err != nil {
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
	var updateData repositories.BarangUpdateRequest
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	repo := repositories.NewBarangRepository()
	barang, err := repo.UpdateBarang(id, &updateData)
	if err != nil {
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
	var updateData repositories.PerusahaanUpdateRequest
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
		})
		return
	}

	repo := repositories.NewPerusahaanRepository()
	company, err := repo.UpdatePerusahaan(id, &updateData)
	if err != nil {
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
