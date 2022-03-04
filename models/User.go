package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserName  string             `json:"username"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Authtoken string             `json:"auth_token"`
	Role      Role
	Company   Company
}
