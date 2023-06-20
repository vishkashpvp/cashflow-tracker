package mongodb

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers() ([]models.User, error) {
	usersColl := GetUsersCollection()

	filter := bson.M{}
	options := options.Find()

	cursor, err := usersColl.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}

func CreateUser(user *models.User) (primitive.ObjectID, int, error) {
	usersColl := GetUsersCollection()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectID()

	resp, err := usersColl.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, http.StatusInternalServerError, err
	}

	insertedId, ok := resp.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, http.StatusBadRequest, errors.New("invalid inserted id")
	}

	return insertedId, http.StatusCreated, nil
}

func FindUserByEmail(email string) (*models.User, int, error) {
	usersColl := GetUsersCollection()
	filter := bson.M{"email": email}

	var user models.User
	err := usersColl.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func FindUserByProviderID(provider, provider_id string) (*models.User, int, error) {
	usersColl := GetUsersCollection()
	filter := bson.M{"provider": provider, "provider_id": provider_id}

	var user models.User
	err := usersColl.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func IsEmailExists(email string) (bool, error) {
	usersColl := GetUsersCollection()
	filter := bson.M{"email": email}

	count, err := usersColl.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
