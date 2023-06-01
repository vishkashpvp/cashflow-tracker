package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	client, err := ConnectToMongoDB()
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

	r.Run(":8080")
}
