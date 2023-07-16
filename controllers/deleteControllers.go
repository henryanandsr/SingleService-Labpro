package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	var company model.Company
	if err := initializers.DB.Where("id = ?", id).First(&company).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Company not found",
		})
		return
	}

	if err := initializers.DB.Delete(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Company deleted successfully",
		"data":    company,
	})
}
