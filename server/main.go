package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/middleware"
	"github.com/vishkashpvp/cashflow-tracker/server/routes"
	"github.com/vishkashpvp/cashflow-tracker/server/utils"
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

	r.Use(middleware.MongoDBClientCheck(client))

	defer client.Disconnect(context.Background())

	routes.Ping(r)
	routes.Auth(r)
	routes.Users(r)

	r.Run(":8080")
}
