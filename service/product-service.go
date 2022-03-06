package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
	m "webappsapi/main/models"
)

var productCollection *mongo.Collection

func init() {
	db := config.Connect()
	productCollection = db.Collection("product")
}

func InsertOneProduct(product []byte) (*mongo.InsertOneResult, error) {
	result_data, err := productCollection.InsertOne(config.Ctx, product)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}

func FindProductById(id string) m.Product {
	result := m.Product{}
	productId, err := primitive.ObjectIDFromHex(id)
	fmt.Println(productId)
	if err != nil {
		log.Println(err)
		return result
	}
	err = productCollection.FindOne(context.TODO(), bson.M{"_id": productId}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func GetProductList() []m.Product {
	results := []m.Product{}
	cursor, err := productCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var product m.Product
		cursor.Decode(&product)
		results = append(results, product)
	}
	return results
}
