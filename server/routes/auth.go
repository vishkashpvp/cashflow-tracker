package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/handlers"
)

func Auth(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/signin", handlers.SignIn)
}
