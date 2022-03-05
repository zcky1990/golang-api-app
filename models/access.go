package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Access struct {
	Id          primitive.ObjectID `bson:"_id"`
	AccessLevel int64              `bson:"access_level" json:"access_level"`
	ListUrl     []string           `bson:"list_url" json:"list_url"`
}
