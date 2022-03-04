package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessLevel struct {
	ApiLevel    int64   `json:"api_level"`
	Description string  `json:"description"`
	AccessLevel []int64 `json:"access_level"`
}

type Role struct {
	Id          primitive.ObjectID `bson:"_id"`
	RoleName    string             `json:"role_name"`
	Description string             `json:"description"`
	Access      []AccessLevel      `json:"access"`
}
