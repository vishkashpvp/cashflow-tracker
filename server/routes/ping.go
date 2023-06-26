package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(r *gin.Engine) {
	r.GET("/ping", pong)
}

func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
