package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoDBClientCheck(client *mongo.Client) gin.HandlerFunc {
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
