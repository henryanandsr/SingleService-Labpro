package main

import (
	"SingleService-Labpro/controllers"
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables(".env")
	initializers.ConnectToDB()
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// print header
		fmt.Println("Received header: ", c.Request.Header)
		tokenString := c.GetHeader("Authorization")

		fmt.Println("Received token string: ", tokenString)

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			return
		}

		claims := &jwt.StandardClaims{}
		tokenParts := strings.Split(tokenString, " ")

		var jwtTokenPart string
		if len(tokenParts) >= 2 {
			jwtTokenPart = tokenParts[1]
		} else {
			jwtTokenPart = tokenParts[0]
		}

		token, err := jwt.ParseWithClaims(jwtTokenPart, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.Jwtkey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("Signature invalid for token: ", tokenString)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				return
			}
			fmt.Println("Error parsing token: ", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !token.Valid {
			fmt.Println("Invalid token: ", tokenString)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://single-service-labpro.vercel.app", "https://monolith-full-stack.vercel.app", "https://ohl-fe.vercel.app", "http://localhost:5173", "http://localhost:8001"}
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
	log.Println("User created")
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
	r.POST("/login", controllers.Login)
	r.PUT("/barang/:id", controllers.UpdateBarang)
	r.PUT("/barang/stok/:id", controllers.UpdateStokBarang)
	r.GET("/barang", controllers.GetBarangs)
	r.GET("/barang/:id", controllers.GetBarang)

	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
