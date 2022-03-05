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

var accessCollection *mongo.Collection

func init() {
	db := config.Connect()
	accessCollection = db.Collection("access")
}

func InsertOneAccess(access []byte) (*mongo.InsertOneResult, error) {
	result_data, err := accessCollection.InsertOne(config.Ctx, access)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}

func FindAccessByAccessLevel(api_level int64) m.Access {
	result := m.Access{}
	err := accessCollection.FindOne(context.TODO(), bson.M{"access_level": api_level}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func FindAccessById(id string) m.Access {
	result := m.Access{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return result
	}
	err = accessCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func FindAccessByIdAndUrl(id string, path string) m.Access {
	result := m.Access{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return result
	}
	err = accessCollection.FindOne(context.TODO(), bson.M{"_id": oid, "list_url": path}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}
