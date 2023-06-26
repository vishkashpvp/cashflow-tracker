package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/routes"
	"github.com/vishkashpvp/cashflow-tracker/server/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	corsConfig := utils.GetCorsConfig()
	r.Use(cors.New(corsConfig))

	client, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Println("Connection to MongoDB failed:", err)
	}

	r.Use(mongodbMiddleware(client))

	defer client.Disconnect(context.Background())

	routes.Ping(r)
	routes.Auth(r)
	routes.User(r)

	r.Run(":8080")
}

func mongodbMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if client == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to the database"})
			c.Abort()
			return
		}

		c.Set("mongoClient", client)
		c.Next()
	}
}
