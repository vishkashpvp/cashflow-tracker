package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {
	users, err := mongodb.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserProfile(c *gin.Context) {
	userMap, _ := c.Get("user")
	user, _ := userMap.(*models.User)
	c.Redirect(http.StatusTemporaryRedirect, "/users/"+user.ID.Hex())
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, statusCode, err := mongodb.FindUserByID(objectID)
	if err != nil {
		c.JSON(statusCode, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
