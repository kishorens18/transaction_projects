package transactionservice

import (
	"context"
	"encoding/json"
	"fmt"
	"transactions/config"
	models "transactions/tran_models"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TransactionContext() *mongo.Collection {
	client, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return config.GetCollection(client, "sample_analytics", "transactions")
}

func FindProducts() ([]*models.Transactionses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"transaction_count": bson.D{{
		Key: "$gt", Value: 85},
		{Key: "$lt", Value: 90},
	}}

	client, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	collection := config.GetCollection(client, "inventory", "students")

	// options := options.Find().SetSort(bson.D{{Key: "transaction_count", Value: -1}}).SetSkip(30).SetLimit(1) // You can configure options here if needed
	options := options.Find()
	result, err := collection.Find(ctx, filter, options)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer result.Close(ctx)

	var products []*models.Transactionses
	for result.Next(ctx) {
		product := &models.Transactionses{}
		err := result.Decode(product)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		//fmt.Println(product)

		products = append(products, product)

	}

	if err := result.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func FetchAggregate() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	matchstage := bson.D{{Key: "$match", Value: bson.D{{Key: "transaction_count", Value: 100}}}}

	groupsatge := bson.D{
		{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$account_id"},
				{Key: "total_count", Value: bson.D{{Key: "$sum", Value: "$transaction_count"}}},
			}}}
	result, err := TransactionContext().Aggregate(ctx, mongo.Pipeline{matchstage, groupsatge})
	if err != nil {
		fmt.Println(err.Error())

	} else {

		var showswithinfo []bson.M
		if err = result.All(ctx, &showswithinfo); err != nil {
			panic(err)
		}
		formatted_data, err := json.MarshalIndent(showswithinfo, "", " ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(formatted_data))
		}
	}
}

func FetchUnwind() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	unwindStage := bson.D{{Key: "$unwind", Value: "$course"}}

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "course", Value: "Python"}}}}

	pipeline := mongo.Pipeline{unwindStage, matchStage}

	result, err := TransactionContext().Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var studentsWithPython []bson.M
		if err = result.All(ctx, &studentsWithPython); err != nil {
			panic(err)
		}
		formattedData, err := json.MarshalIndent(studentsWithPython, "", " ")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(formattedData))
		}
	}
}

func UpdateTransaction(initialvalue int, newValue int) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "account_id", Value: initialvalue}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "account_id", Value: newValue}}}}
	result, err := TransactionContext().UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil

}
