package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"google.golang.org/api/idtoken"
)

// SignIn validates the idToken based on the provider('X-Provider' header) and constructs a response with the user information.
// If any error occurs during validation or construction, it returns an appropriate error response.
func SignIn(c *gin.Context) {
	idToken := c.Request.Header.Get("X-IdToken")
	provider := strings.ToUpper(c.Request.Header.Get("X-Provider"))

	switch provider {
	case "GOOGLE":
		user, err := GoogleUserInfo(idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown provider '" + provider + "'"})
	}
}

// GoogleUserInfo validates idToken and retrieves user info from provided idToken.
// It returns a *models.User and an error, if any.
func GoogleUserInfo(idToken string) (*models.User, error) {
	clientId := os.Getenv("G_CLOUD_CLIENT_ID")
	if clientId == "" {
		return nil, errors.New("client Id is missing")
	}

	payload, err := idtoken.Validate(context.Background(), idToken, clientId)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    payload.Claims["email"].(string),
		Name:     payload.Claims["name"].(string),
		Provider: "GOOGLE",
		Picture:  payload.Claims["picture"].(string),
	}

	return user, nil
}
