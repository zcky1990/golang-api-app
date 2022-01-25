package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
	"webappsapi/main/jwtconfig"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
}

type UserLogin struct {
	Id            primitive.ObjectID `bson:"_id"`
	Username      string             `json:"username"`
	Email         string             `json:"email"`
	Firstname     string             `json:"firstname"`
	Lastname      string             `json:"lastname"`
	Authorization string             `json:"auth_token"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UserAddRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
}

type UserSuccessResponse struct {
	Status string    `json:"status"`
	Code   uint8     `json:"code"`
	Data   UserLogin `json:"data"`
}

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

func GetUserByEmailAndPassword(email string, password string) (UserLogin, error) {
	result := UserLogin{}
	// var podcast bson.M
	// collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&podcast)
	// fmt.Println("podcast : ", podcast)
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
		response := GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)

	}
	token, err := jwtconfig.CreateToken(user.Id, user.Email)
	if err != nil {
		response := GetFailedResponse(err.Error())
		json.NewEncoder(w).Encode(response)
	}
	user.Authorization = token
	response := UserSuccessResponse{"success", 200, user}
	json.NewEncoder(w).Encode(response)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request UserAddRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// email := request.Email
	// password := request.Password
	// firstname := request.Firstname
	// lastname := request.Lastname
	// birthday := request.Birthday
	// username := request.Username

	// user_data := bson.D{
	// 	{Key: "email", Value: email},
	// 	{Key: "password", Value: password},
	// 	{Key: "firstname", Value: firstname},
	// 	{Key: "lastname", Value: lastname},
	// 	{Key: "birthday", Value: birthday},
	// 	{Key: "username", Value: username},
	// }
	user := GetUserByEmail(request.Email)
	if (user != User{}) {
		json.NewEncoder(w).Encode(GetFailedResponse("User exists"))
	} else {
		//convert struct to bson
		user_data, err := bson.Marshal(request)
		if err != nil {
			panic(err)
		}

		user, err := addUser(user_data)
		if err != nil {
			response := GetFailedResponse(err.Error())
			json.NewEncoder(w).Encode(response)
		}
		id := user.InsertedID.(primitive.ObjectID).Hex()
		response := GetSuccessResponse()
		response.Data["id"] = id
		response.Data["email"] = request.Email
		response.Data["firstname"] = request.Firstname
		response.Data["lastname"] = request.Lastname
		json.NewEncoder(w).Encode(response)
	}
}
