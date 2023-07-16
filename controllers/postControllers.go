package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostBarang(c *gin.Context) {
	var request struct {
		NamaBarang   string `json:"nama"`
		HargaBarang  int    `json:"harga"`
		StokBarang   int    `json:"stok"`
		PerusahaanID string `json:"perusahaan_id"`
		KodeBarang   string `json:"kodeBarang"`
	}
	// if request.HargaBarang <= 0 || request.StokBarang < 0 || request.PerusahaanID == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "error",
	// 		"message": "Invalid data",
	// 		"data":    nil,
	// 	})
	// 	return
	// }
	var existingCompany model.Company
	if err := initializers.DB.Where("id = ?", request.PerusahaanID).First(&existingCompany).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Perusahaan not found",
			"data":    nil,
		})
		return
	}
	var existingBarang model.Barang
	if !gorm.IsRecordNotFoundError(initializers.DB.Where("kode_barang = ?", request.KodeBarang).First(&existingBarang).Error) {
		c.JSON(http.StatusConflict, gin.H{
			"status":  "error",
			"message": "KodeBarang already exists",
			"data":    nil,
		})
		return
	}
	var highestIDBarang model.Barang
	initializers.DB.Order("kode_barang desc").First(&highestIDBarang)
	highestID, _ := strconv.Atoi(highestIDBarang.KodeBarang)
	newID := highestID + 1
	newIDString := strconv.Itoa(newID)

	barang := &model.Barang{
		KodeBarang:   newIDString,
		NamaBarang:   request.NamaBarang,
		HargaBarang:  request.HargaBarang,
		StokBarang:   request.StokBarang,
		IDPerusahaan: existingCompany.Nama,
	}
	initializers.DB.Create(barang)
}

func PostCompany(c *gin.Context) {
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
	if len(request.Kode) != 3 || strings.ToUpper(request.Kode) != request.Kode {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid Kode, it must be all upper case and have a length of 3",
		})
		return
	}
	var highestIDCompany model.Company
	initializers.DB.Order("id desc").First(&highestIDCompany)
	highestID, _ := strconv.Atoi(highestIDCompany.ID)
	newID := highestID + 1
	newIDString := strconv.Itoa(newID)
	company := &model.Company{
		ID:        newIDString,
		Nama:      request.Nama,
		Alamat:    request.Alamat,
		NoTelepon: request.NoTelp,
		KodePajak: request.Kode,
	}
	result := initializers.DB.Create(company)
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
