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
