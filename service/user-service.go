package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
	m "webappsapi/main/models"
)

type User = m.User

var collection *mongo.Collection

func init() {
	db := config.Connect()
	collection = db.Collection("Users")
}

func AddUser(user []byte) (*mongo.InsertOneResult, error) {
	result_data, err := collection.InsertOne(config.Ctx, user)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return result_data, nil
}

func GetUserByEmail(email string) User {
	result := User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func GetUserByEmailAndPassword(email string, password string) (User, error) {
	result := User{}
	err := collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result, err
	}
	return result, nil
}

func GetAllUserList() []User {
	results := []User{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		results = append(results, user)
	}
	return results
}

func SearchUser(search_type string, query string) []User {
	results := []User{}
	filter := bson.M{}
	if search_type == "name" {
		filter = bson.M{"$or": []interface{}{
			bson.M{"firstName": bson.M{"$regex": query, "$options": "im"}},
			bson.M{"lastName": bson.M{"$regex": query, "$options": "im"}},
		},
		}
	}

	if search_type == "email" {
		filter = bson.M{"email": query}
	}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		return results
	}

	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		results = append(results, user)
	}
	fmt.Println(results)
	return results
}
