package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var request struct {
		NamaBarang   string `json:"nama"`
		HargaBarang  int    `json:"harga"`
		StokBarang   int    `json:"stok"`
		PerusahaanID string `json:"perusahaan_id"`
		KodeBarang   string `json:"kode"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid JSON format",
			"data":    nil,
		})
		return
	}

	db, err := initializers.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	existingBarang := &model.Barang{}
	result := db.Where("kode_barang = ?", request.KodeBarang).First(existingBarang)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "error",
			"message": "Barang with the same KodeBarang already exists",
		})
		return
	}

	barang := &model.Barang{
		ID:                uuid.New().String(),
		KodeBarang:        request.KodeBarang,
		NamaBarang:        request.NamaBarang,
		HargaBarang:       request.HargaBarang,
		StokBarang:        request.StokBarang,
		PerusahaanPembuat: request.PerusahaanID,
	}
	result = db.Create(barang)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
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
	var request struct {
		Nama   string `json:"nama"`
		Alamat string `json:"alamat"`
		NoTelp string `json:"no_telp"`
		Kode   string `json:"kode"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid data format",
		})
		return
	}
	db, err := initializers.GetDBInstance() // added this line
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	company := &model.Company{
		ID:        uuid.New().String(),
		Nama:      request.Nama,
		Alamat:    request.Alamat,
		NoTelepon: request.NoTelp,
		KodePajak: request.Kode,
	}

	result := db.Create(company)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Company created successfully",
		"data":    company,
	})
}
