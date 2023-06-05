package mongodb

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllUsers retrieves all users from the db.
//
// It returns a slice of 'models.User' struct containing all the users found in the database or an error if any.
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

// CreateUser creates a new user in the db with the provided user data.
//
// It takes a pointer to a 'models.User' struct as input and assigns a new ObjectID as the user's ID.
// The function checks if the email already exists in the db, if it exists, returns an error with an HTTP status code of 409 (Conflict).
// Otherwise, it inserts the user into the db and returns the inserted ID along with an HTTP status code of 201 (Created).
func CreateUser(user *models.User) (primitive.ObjectID, int, error) {
	usersColl := GetUsersCollection()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.ID = primitive.NewObjectID()

	emailExists, err := IsEmailExists(user.Email)
	if err != nil {
		return primitive.NilObjectID, http.StatusInternalServerError, err
	}
	if emailExists {
		return primitive.NilObjectID, http.StatusConflict, errors.New("email already exists")
	}

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

// IsEmailExists checks if an email already exists in the database.
// It returns true if the email exists, false if it doesn't,
// or an error if an error occurs during the database query.
func IsEmailExists(email string) (bool, error) {
	usersColl := GetUsersCollection()
	filter := bson.M{"email": email}

	count, err := usersColl.CountDocuments(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
