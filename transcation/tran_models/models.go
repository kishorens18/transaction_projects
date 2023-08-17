package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Date            time.Time `json:"date" bson:"date"`
	Amount          int       `json:"amount" bson:"amount"`
	TransactionCode string    `json:"transaction_code" bson:"transaction_code"`
	Symbol          string    `json:"symbol" bson:"symbol"`
	Price           string    `json:"price" bson:"price"`
	Total           string    `json:"total" bson:"total"`
}

type Transactionses struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	AccountID         int                `json:"account_id" bson:"account_id"`
	Transaction_count int                `json:"transaction_count" bson:"transaction_count"`
	Bucket_startdate  time.Time          `json:"bucket_start_date" bson:"bucket_start_date"`
	Bucket_enddate    time.Time          `json:"bucket_end_date" bson:"bucket_end_date"`
	// Transactions   []Transaction `json:"transactions" bson:"transactions"`
}
