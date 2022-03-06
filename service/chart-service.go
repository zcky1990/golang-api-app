package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
	m "webappsapi/main/models"
)

var chartCollection *mongo.Collection

func init() {
	db := config.Connect()
	chartCollection = db.Collection("chart")
}

func InsertOneChart(chart []byte) (*mongo.InsertOneResult, error) {
	result_data, err := chartCollection.InsertOne(config.Ctx, chart)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}

func GetChartListByUserId(user_id string) []m.ChartList {
	results := []m.ChartList{}
	userId, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		log.Println(err)
		return results
	}

	filter := []bson.M{{"$lookup": bson.M{"from": "product", "localField": "product_id", "foreignField": "_id", "as": "ps"}}, {"$unwind": bson.M{"path": "$ps"}}, {"$match": bson.M{"user_id": userId}}, {"$project": bson.M{"id": "$_id", "product_id": "$ps._id", "title": "$ps.title", "price": "$ps.price", "picture_url": "$ps.picture_url"}}}
	cursor, err := chartCollection.Aggregate(context.TODO(), filter)
	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var chart m.ChartList
		cursor.Decode(&chart)
		results = append(results, chart)
	}
	return results
}
