package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"username"  json:"username"`
	Email     string             `bson:"email"  json:"email"`
	FirstName string             `bson:"firstname"  json:"firstname"`
	LastName  string             `bson:"lastname"  json:"lastname"`
	Authtoken string             `bson:"auth_token"  json:"auth_token"`
	RoleId    primitive.ObjectID `bson:"role_id"  json:"role_id"`
	CompanyId primitive.ObjectID `bson:"company_id"  json:"company_id"`
}

type InsertUserMongoModels struct {
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Firstname string             `bson:"firstname" json:"firstname"`
	Lastname  string             `bson:"lastname" json:"lastname"`
	Password  string             `bson:"password" json:"password"`
	Birthday  string             `bson:"birthday" json:"birthday"`
	RoleId    primitive.ObjectID `bson:"role_id" json:"role_id"`
	CompanyId primitive.ObjectID `bson:"company_id" json:"company_id"`
}
