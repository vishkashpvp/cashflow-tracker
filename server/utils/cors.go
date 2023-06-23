package utils

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
)

func GetCorsConfig() cors.Config {
	allowedOriginsStr := os.Getenv("ALLOWED_ORIGINS")
	allowedOrigins := strings.Split(allowedOriginsStr, ",")

	config := cors.DefaultConfig()
	config.AllowOrigins = allowedOrigins
	config.AllowHeaders = append(config.AllowHeaders, "x-idtoken", "x-provider", "x-accesstoken")

	return config
}
