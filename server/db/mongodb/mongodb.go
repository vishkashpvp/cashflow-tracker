package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const LOCAL_MONGO string = "mongodb://localhost:27017"

var client *mongo.Client

func ConnectToMongoDB() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = LOCAL_MONGO
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client with the client options
	newClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Establish a connection to the MongoDB server
	err = newClient.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to verify the connection
	err = newClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	client = newClient
	log.Println("Connected to MongoDB")

	return client, nil
}

func GetAppDatabase() *mongo.Database {
	dbName := os.Getenv("MONGODB_APPDB")
	return client.Database(dbName)
}

func GetUsersCollection() *mongo.Collection {
	return GetAppDatabase().Collection("users")
}
