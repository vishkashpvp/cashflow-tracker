package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/handlers"
)

func User(r *gin.Engine) {
	user := r.Group("/user")

	user.GET("/all", handlers.GetAllUsers)
	user.GET(":id", handlers.GetUserByID)
}
