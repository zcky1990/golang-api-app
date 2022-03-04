package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id          primitive.ObjectID `bson:"_id"`
	Rolename    string             `json:"rolename"`
	Description string             `json:"description"`
}
