package controllers

import (
	repositories "SingleService-Labpro/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostBarang(c *gin.Context) {
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
	var request repositories.BarangPostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON format",
			"data":    nil,
		})
		return
	}

	repo := repositories.NewBarangRepository()
	barang, err := repo.CreateBarang(&request)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Barang created successfully",
		"data":    barang,
	})
}

func PostCompany(c *gin.Context) {
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
	var request repositories.PerusahaanPostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data format",
		})
		return
	}

	repo := repositories.NewPerusahaanRepository()
	company, err := repo.CreatePerusahaan(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Company created successfully",
		"data":    company,
	})
}
