package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Authtoken string             `json:"auth_token"`
	Role      Role
	Company   Company
}
