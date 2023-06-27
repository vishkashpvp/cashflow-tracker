package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/handlers"
)

func Users(r *gin.Engine) {
	users := r.Group("/users")

	users.GET("/all", handlers.GetAllUsers)
	users.GET(":id", handlers.GetUserByID)
}
