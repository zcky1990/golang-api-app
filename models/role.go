package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id            primitive.ObjectID `bson:"_id"`
	RoleName      string             `bson:"role_name" json:"role_name"`
	Description   string             `bson:"description" json:"description"`
	AccessLevelId primitive.ObjectID `bson:"access_level_id" json:"access_level_id"`
}

type InsertRoleMongoModels struct {
	RoleName      string             `bson:"role_name" json:"role_name"`
	Description   string             `bson:"description" json:"description"`
	AccessLevelId primitive.ObjectID `bson:"access_level_id" json:"access_level_id"`
}
