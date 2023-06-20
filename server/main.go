package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(getCorsConfig()))

	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Println("Connection to MongoDB failed:", err)
	}

	r.Use(func(c *gin.Context) {
		if client == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
			c.Abort()
			return
		}

		c.Set("mongoClient", client)
		c.Next()
	})

	defer client.Disconnect(context.Background())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Go!"})
	})
	r.POST("/signin", handlers.SignIn)
	r.GET("/user/all", handlers.GetAllUsers)

	r.Run(":8080")
}

func getCorsConfig() cors.Config {
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(allowedOriginsStr, ",")

	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	config.AllowHeaders = append(config.AllowHeaders, "x-idtoken", "x-provider", "x-accesstoken")

	return config
}
