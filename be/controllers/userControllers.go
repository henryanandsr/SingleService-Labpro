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

var Jwtkey = []byte("gfdsgfdsgfdsgfdsgdf")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c *gin.Context) {
	fmt.Println("1")
	var credentials model.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format", "data": nil})
		return
	}
	fmt.Println("2")
	fmt.Println("Received credentials: ", credentials)
	var User model.User
	if err := initializers.DB.Where("username = ?", credentials.Username).First(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid username", "data": nil})
		return
	}
	fmt.Println("3")
	if User.Password != credentials.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid password", "data": nil})
		return
	}
	fmt.Println("4")
	claims := &Claims{
		Username:       credentials.Username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Jwtkey)
	fmt.Println("Token generated:", tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Could not generate token", "data": nil})
		return
	}
	fmt.Println("5")
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
	c.Header("Authorization", "Bearer "+tokenString)
	// print header
	fmt.Println(c.Writer.Header())
}

// Self endpoint
func Self(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.Split(tokenString, " ")[0]
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
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
