package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := extractJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		parsedToken, err := jwt.Parse(token, keyFunc)
		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := parsedToken.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)

		id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		user, statusCode, err := mongodb.FindUserByID(id)
		if err != nil {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func keyFunc(t *jwt.Token) (interface{}, error) {
	secretKeyString := os.Getenv("JWT_SECRET_KEY")
	if secretKeyString == "" {
		return nil, errors.New("secret key not found")
	}
	secretKey := []byte(secretKeyString)
	if t.Method != jwt.SigningMethodHS256 {
		return nil, errors.New("unexpected signing method")
	}
	return secretKey, nil
}

func extractJWT(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format: must be in the format 'Bearer <token>'")
	}
	token := strings.TrimSpace(authParts[1])
	if token == "" {
		return "", errors.New("empty token")
	}
	return token, nil
}
