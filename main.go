package main

import (
	"SingleService-Labpro/controllers"
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"
	"fmt"
	"net/http"
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
		// Get the JWT string from the header
		tokenString := c.GetHeader("Authorization")

		// If the token is empty return an error
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
			return
		}

		// Validate token
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

		// Everything OK, proceed with the request
		c.Next()
	}
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

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
		authorized.GET("/barang/:id", controllers.GetBarang)
		authorized.GET("/perusahaan/:id", controllers.GetPerusahaan)
		authorized.GET("/perusahaan", controllers.GetPerusahaans)
		authorized.GET("/barang", controllers.GetBarangs)
		authorized.PUT("/barang/:id", controllers.UpdateBarang)
		authorized.PUT("/perusahaan/:id", controllers.UpdateCompany)
		authorized.GET("/self", controllers.Self)
	}

	r.POST("/login", controllers.Login)

	r.Run(":8080")
}
