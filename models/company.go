package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id             primitive.ObjectID `bson:"_id"`
	Companyname    string             `json:"company_name"`
	Companyaddress string             `json:"company_address"`
	Companyemail   string             `json:"company_email"`
	Companyphone   string             `json:"company_phone"`
}
