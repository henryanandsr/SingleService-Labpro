package controllers

import (
	model "SingleService-Labpro/models"
	repositories "SingleService-Labpro/repository"
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
	var credentials model.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format", "data": nil})
		return
	}

	repo := repositories.NewUserRepository()
	user, err := repo.FindUserByUsername(credentials.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid username", "data": nil})
		return
	}

	if user.Password != credentials.Password {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid password", "data": nil})
		return
	}

	claims := &Claims{
		Username:       credentials.Username,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Jwtkey)
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
	c.Header("Authorization", "Bearer "+tokenString)
}

func Self(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tokenString = strings.Split(tokenString, " ")[1] // Get the token part after "Bearer "
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
