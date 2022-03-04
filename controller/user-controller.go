package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"

	"webappsapi/main/config"
	"webappsapi/main/jwtconfig"
	m "webappsapi/main/models"
	rq "webappsapi/main/request"
	rs "webappsapi/main/response"
)

type UserLoginRequest = rq.UserLoginRequest
type UserSignUpRequest = rq.UserSignUpRequest
type UserAddRequest = rq.UserAddRequest
type User = m.User

var collection *mongo.Collection

func init() {
	db := config.Connect()
	collection = db.Collection("Users")
}

func addUser(user []byte) (*mongo.InsertOneResult, error) {
	result_data, err := collection.InsertOne(ctx, user)
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

func UserList(w http.ResponseWriter, r *http.Request) {
	response := GetAllUserList()
	json.NewEncoder(w).Encode(response)
}

func Search(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	query := values["query"][0]
	search_type := values["search_type"][0]

	response := SearchUser(search_type, query)
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var request UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := request.Email
	password := request.Password
	user, err := GetUserByEmailAndPassword(email, password)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	token, err := jwtconfig.CreateToken(user.Id, user.Email)
	if err != nil {
		response := rs.GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	user.Authtoken = token
	response := rs.GetSuccessResponse(&fiber.Map{"data": user})
	json.NewEncoder(w).Encode(response)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request UserAddRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := GetUserByEmail(request.Email)
	if (user != User{}) {
		json.NewEncoder(w).Encode(rs.GetFailedResponse("User exists"))
	} else {
		//convert struct to bson
		user_data, err := bson.Marshal(request)
		if err != nil {
			panic(err)
		}

		user, err := addUser(user_data)
		if err != nil {
			response := rs.GetFailedResponse(err.Error())
			json.NewEncoder(w).Encode(response)
		}
		// id := user.InsertedID.(primitive.ObjectID).Hex()
		response := rs.GetSuccessResponse(&fiber.Map{"data": user})
		json.NewEncoder(w).Encode(response)
	}
}
