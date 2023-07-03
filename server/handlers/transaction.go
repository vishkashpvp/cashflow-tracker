package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTransactions(c *gin.Context) {
	transactions, err := mongodb.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	transaction, statusCode, err := mongodb.FindTransactionByID(objectID)
	if err != nil {
		c.JSON(statusCode, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {
	userMap, _ := c.Get("user")
	transactionMap, _ := c.Get("transaction")

	user, _ := userMap.(*models.User)
	transaction, _ := transactionMap.(*models.Transaction)

	transaction.UserID = user.ID

	createdTransaction, statusCode, err := mongodb.CreateTransaction(transaction)
	if err != nil {
		c.JSON(statusCode, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"transaction": createdTransaction,
		"message":     "Transaction created successfully",
	})
}
