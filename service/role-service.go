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

var roleCollection *mongo.Collection

func init() {
	db := config.Connect()
	roleCollection = db.Collection("role")
}

func FindRoleOwner() m.Role {
	result := m.Role{}
	err := roleCollection.FindOne(context.TODO(), bson.M{"rolename": "owner"}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func FindRoleById(id string) m.Role {
	result := m.Role{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return result
	}
	err = roleCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func FindRoleByName(role_name string) m.Role {
	result := m.Role{}
	err := roleCollection.FindOne(context.TODO(), bson.M{"rolename": role_name}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func InsertOneRole(user []byte) (*mongo.InsertOneResult, error) {
	result_data, err := roleCollection.InsertOne(config.Ctx, user)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}
