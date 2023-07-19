package controllers

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("default_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	var credentials model.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format", "data": nil})
		return
	}

	var User model.User
	if err := initializers.DB.Where("username = ?", credentials.Username).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid username", "data": nil})
		return
	}

	if User.Password != credentials.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid password", "data": nil})
		return
	}

	claims := &Claims{
		Username:       credentials.Username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println("Token generated:", tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Could not generate token", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful",
		"data": gin.H{
			"user": gin.H{
				"username": credentials.Username,
				"name":     "User",
			},
			"token": tokenString,
		},
	})
	c.Writer.Header().Set("Authorization", "Bearer "+tokenString)
}

// Self endpoint
func Self(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.Split(tokenString, " ")[1]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unauthorized", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data retrieved successfully",
		"data": gin.H{
			"username": claims.Username,
			"name":     "User",
		},
	})
}
