package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vishkashpvp/cashflow-tracker/server/db/mongodb"
	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/idtoken"
)

// SignUp handles the user signup functionality.
//
// Based on provider(X-Provider) and idToken(X-IdToken), it calls respective function and constructs user information.
// Then, it creates a new user in the MongoDB.
// It returns the created user and Id in the response if successful, or an appropriate error message if there was an error.
func SignUp(c *gin.Context) {
	idToken := c.Request.Header.Get("X-IdToken")
	provider := strings.ToUpper(c.Request.Header.Get("X-Provider"))

	switch provider {
	case "GOOGLE":
		user, err := GoogleUserInfo(idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		createdId, statusCode, err := mongodb.CreateUser(user)
		if err != nil {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}

		token, err := GenerateJWT(createdId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown provider '" + provider + "'"})
	}
}

// SignIn handles the user signin functionality.
//
// It validates the idToken based on the provider('X-Provider' header) and constructs a response with the user information.
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

		userCopy := *user
		user, statusCode, err := mongodb.FindUserByEmail(user.Email)
		if err != nil && statusCode != http.StatusNotFound {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}
		respStatus := http.StatusOK
		if statusCode == http.StatusNotFound {
			user = &userCopy
			createdId, statusCode, err := mongodb.CreateUser(user)
			if err != nil {
				c.JSON(statusCode, gin.H{"message": err.Error()})
				return
			}

			user.ID = createdId
			respStatus = http.StatusCreated
		}

		token, err := GenerateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(respStatus, gin.H{"token": token, "user": user})
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

// GenerateJWT generates a JSON Web Token (JWT) using the provided user ID.
// It returns jwt token or error if any.
func GenerateJWT(id primitive.ObjectID) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("secret key not found")
	}

	secretKeyBytes := []byte(secretKey)
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"id":  id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
