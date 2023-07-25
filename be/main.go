package main

import (
	"SingleService-Labpro/controllers"
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("default_key")

func init() {
	initializers.LoadEnvVariables(".env")
	initializers.ConnectToDB()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			return
		}

		claims := &jwt.StandardClaims{}
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		token, err := jwt.ParseWithClaims(strings.Split(tokenString, " ")[1], claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://single-service-labpro.vercel.app", "https://monolith-full-stack.vercel.app"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	port := os.Getenv("PORT")

	user := model.User{
		Username: "admin",
		Password: "admin",
	}
	var existingUser model.User
	if err := initializers.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		if err := initializers.DB.Create(&user).Error; err != nil {
			fmt.Println("Could not create user: ", err)
		}
	} else {
		fmt.Println("User already exists")
	}

	authorized := r.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.POST("/barang", controllers.PostBarang)
		authorized.POST("/perusahaan", controllers.PostCompany)
		authorized.DELETE("/perusahaan/:id", controllers.DeleteCompany)
		authorized.DELETE("/barang/:id", controllers.DeleteBarang)
		authorized.GET("/perusahaan/:id", controllers.GetPerusahaan)
		authorized.GET("/perusahaan", controllers.GetPerusahaans)
		authorized.PUT("/perusahaan/:id", controllers.UpdateCompany)
		authorized.GET("/self", controllers.Self)
	}
	r.PUT("/barang/:id", controllers.UpdateBarang)
	r.PUT("/barang/stok/:id", controllers.UpdateStokBarang)
	r.GET("/barang", controllers.GetBarangs)
	r.GET("/barang/:id", controllers.GetBarang)

	r.POST("/login", controllers.Login)
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
