package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TransactionTypeSpending TransactionType = "spending"
	TransactionTypeEarning  TransactionType = "earning"
)

type Transaction struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	UserID        primitive.ObjectID `json:"userID" bson:"user_id"`
	Amount        float64            `json:"amount"`
	Tax           float64            `json:"tax"`
	PaymentMethod string             `json:"paymentMethod" bson:"payment_method"`
	Type          TransactionType    `json:"type"`
	Category      string             `json:"category"`
	Description   string             `json:"description"`
	Date          time.Time          `json:"date"`
	Place         string             `json:"place"`
	Location      Location           `json:"location"`
	Tags          []string           `json:"tags"`
	CreatedAt     time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updated_at"`
}

type Location struct {
	Area    string `json:"area"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type TransactionType string

func (tt TransactionType) IsValid() bool {
	return tt == TransactionTypeSpending || tt == TransactionTypeEarning
}
