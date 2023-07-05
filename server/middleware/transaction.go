package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"github.com/vishkashpvp/cashflow-tracker/server/transform"
)

func BindAndValidateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		transaction := &models.Transaction{
			Tags: []string{},
		}

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

		category, err := transform.TransformCategory(transaction.Category)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		transaction.Category = category

		tt, err := transform.TransformTransactionType(transaction.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		transaction.Type = tt

		if transaction.Date.IsZero() {
			transaction.Date = time.Now().UTC()
		}

		c.Set("transaction", transaction)
		c.Next()
	}
}
