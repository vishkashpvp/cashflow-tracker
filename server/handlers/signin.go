package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
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
	accessToken := c.Request.Header.Get("X-AccessToken")

	user, err := getUserInfo(provider, idToken, accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	userCopy := *user

	user, statusCode, err := mongodb.FindUserByProviderID(provider, user.Provider_ID)
	if err != nil && statusCode != http.StatusNotFound {
		c.JSON(statusCode, gin.H{"message": err.Error()})
		return
	}
	respStatus := http.StatusOK
	if statusCode == http.StatusNotFound {
		user = &userCopy
		createdID, statusCode, err := mongodb.CreateUser(user)
		if err != nil {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}

		user.ID = createdID
		respStatus = http.StatusCreated
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(respStatus, gin.H{"token": token, "user": user})
}

func getUserInfo(provider, idToken, accessToken string) (*models.User, error) {
	switch provider {
	case "GOOGLE":
		return GoogleUserInfo(idToken)
	case "FACEBOOK":
		return FacebookUserInfo(accessToken)
	default:
		return nil, errors.New("no such provider: " + provider)
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
		Provider_ID: payload.Subject,
		Email:       payload.Claims["email"].(string),
		Name:        payload.Claims["name"].(string),
		Provider:    "GOOGLE",
		Picture:     payload.Claims["picture"].(string),
	}

	return user, nil
}

func FacebookUserInfo(accessToken string) (*models.User, error) {
	userInfoURL := "https://graph.facebook.com/me?access_token=" + accessToken + "&fields=id,name,email,picture"

	resp, err := http.Get(userInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fbUser models.FacebookUser

	err = json.Unmarshal(body, &fbUser)
	if err != nil {
		return nil, err
	}
	if fbUser.ID == "" {
		return nil, errors.New("facebook user id is missing for provided token")
	}

	user := &models.User{
		Provider_ID: fbUser.ID,
		Email:       fbUser.Email,
		Name:        fbUser.Name,
		Provider:    "FACEBOOK",
		Picture:     fbUser.Picture.Data.URL,
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
