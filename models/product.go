package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id         primitive.ObjectID `bson:"_id"`
	Tilte      string             `bson:"title" json:"title"`
	PictureUrl string             `bson:"picture_url" json:"picture_url"`
	Price      float32            `bson:"price" json:"price"`
}

type InsertProductMongoModels struct {
	Tilte      string  `bson:"title" json:"title"`
	PictureUrl string  `bson:"picture_url" json:"picture_url"`
	Price      float32 `bson:"price" json:"price"`
}
