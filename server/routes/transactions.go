package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/handlers"
	"github.com/vishkashpvp/cashflow-tracker/server/middleware"
)

func Transactions(r *gin.Engine) {
	transactions := r.Group("/transactions")

	transactions.GET("/all", handlers.GetAllTransactions)
	transactions.GET("/:id", handlers.GetTransactionByID)
	transactions.POST("/", middleware.BindAndValidateTransaction(), handlers.CreateTransaction)
}
