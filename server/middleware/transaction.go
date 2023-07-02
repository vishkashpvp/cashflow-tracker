package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
)

func BindAndValidateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		transaction := &models.Transaction{}

		if err := c.ShouldBindJSON(transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
			c.Abort()
			return
		}

		if transaction.Amount <= 0 {
			errMsg := "Invalid amount: " + strconv.FormatFloat(transaction.Amount, 'f', -1, 64)
			c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
			c.Abort()
			return
		}
		if transaction.Category == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
			c.Abort()
			return
		}
		if !transaction.Type.IsValid() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction type must be either 'earning' or 'spending'"})
			c.Abort()
			return
		}

		if transaction.Date.IsZero() {
			transaction.Date = time.Now().UTC()
		}

		c.Set("transaction", transaction)
		c.Next()
	}
}
