package main

import (
	"fmt"
	transactionservice "transactions/transcations"

	//"mongodb/restaurantproducts"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
)

func main() {
	fmt.Println("MongoDB successfully connected...")

	// products := []interface{}{
	//  models.Product{ID: primitive.NewObjectID(), Name: "OnePlus", Price: 1000000, Description: "Budget Phone"},
	//  models.Product{ID: primitive.NewObjectID(), Name: "Vivo", Price: 100000, Description: "China based Phone"},
	// }

	// services.InsertProductList(products)

	// products, _ := transactionservice.FindProducts()
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// transactionservice.FetchUnwind()

	result, err := transactionservice.UpdateTransaction(4431789, 4431787)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result.ModifiedCount)
	}
}

//to connect to data base
//we need to have database name in that database which collection these are the 2 things required
//creating the modules (schema)

//optins .client .uri offered by mongo db connection string to go formate we are converting
//it is offering you a method called connect, also disconnect also will be offered by the instance of it

// ************          model (schema)      *****************

//binging: reqired
//bson :from and to database collection name
// var prdductcollection *mongo.collection=config.Getcollection(config.DB,"products")
// products-column name
// let us create a common method that take two parameters the first is database name and second collection name\
//                and this method will return mongo colleciton
//once you get he reference of the collection you can perform any query using it(crud)

//
