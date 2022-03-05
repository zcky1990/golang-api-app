package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id             primitive.ObjectID `bson:"_id"`
	CompanyName    string             `bson:"company_name" json:"company_name"`
	CompanyAddress string             `bson:"company_address" json:"company_address"`
	CompanyEmail   string             `bson:"company_email" json:"company_email"`
	CompanyPhone   string             `bson:"company_phone" json:"company_phone"`
}
