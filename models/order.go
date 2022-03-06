package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id        primitive.ObjectID `bson:"_id"`
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
	UserId    primitive.ObjectID `bson:"user_Id" json:"user_Id"`
}

type InsertOrderMongoModels struct {
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
	UserId    primitive.ObjectID `bson:"user_Id" json:"user_Id"`
}
