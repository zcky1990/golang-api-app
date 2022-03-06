package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chart struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
}

type ChartList struct {
	Id         primitive.ObjectID `bson:"_id"`
	ProductId  primitive.ObjectID `bson:"product_id" json:"product_id"`
	Tilte      string             `bson:"title" json:"title"`
	PictureUrl string             `bson:"picture_url" json:"picture_url"`
	Price      float32            `bson:"price" json:"price"`
}

type ChartListWithUserId struct {
	Id         primitive.ObjectID `bson:"_id"`
	Tilte      string             `bson:"title" json:"title"`
	PictureUrl string             `bson:"picture_url" json:"picture_url"`
	Price      float32            `bson:"price" json:"price"`
	UserId     primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type InsertChartMongoModels struct {
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductId primitive.ObjectID `bson:"product_id" json:"product_id"`
}
