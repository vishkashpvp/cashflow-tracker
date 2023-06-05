package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
)

func GetAllUsers(c *gin.Context) {
	users, err := mongodb.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
