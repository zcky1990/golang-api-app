package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id             primitive.ObjectID `bson:"_id"`
	CompanyName    string             `json:"company_name"`
	CompanyAddress string             `json:"company_address"`
	CompanyEmail   string             `json:"company_email"`
	CompanyPhone   string             `json:"company_phone"`
}
