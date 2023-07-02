package mongodb

import (
	"context"
	"net/http"
	"time"

	"github.com/vishkashpvp/cashflow-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTransactions() ([]models.Transaction, error) {
	transactionsColl := GetTransactionsCollection()

	filter := bson.M{}
	options := options.Find()

	cursor, err := transactionsColl.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var transactions []models.Transaction
	for cursor.Next(context.Background()) {
		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return []models.Transaction{}, nil
	}

	return transactions, nil
}

func CreateTransaction(transaction *models.Transaction) (*models.Transaction, int, error) {
	transactionColl := GetTransactionsCollection()

	transaction.ID = primitive.NewObjectID()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	_, err := transactionColl.InsertOne(context.Background(), transaction)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return transaction, http.StatusCreated, nil
}
