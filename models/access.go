package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Access struct {
	Id          primitive.ObjectID `bson:"_id"`
	AccessLevel int64              `bson:"access_level" json:"access_level"`
	AccessType  string             `bson:"access_type" json:"access_type"`
	ListUrl     []string           `bson:"list_url" json:"list_url"`
}

type InsertAccessMongoModelsWithCompanyId struct {
	AccessLevel int64              `bson:"access_level" json:"access_level"`
	AccessType  string             `bson:"access_type" json:"access_type"`
	ListUrl     []string           `bson:"list_url" json:"list_url"`
	CompanyId   primitive.ObjectID `bson:"company_id" json:"company_id"`
}
