package service

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
)

var orderCollection *mongo.Collection

func init() {
	db := config.Connect()
	orderCollection = db.Collection("order")
}

func InsertOneOrder(order []byte) (*mongo.InsertOneResult, error) {
	result_data, err := orderCollection.InsertOne(config.Ctx, order)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}
