package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"webappsapi/main/config"
	m "webappsapi/main/models"
)

var companyCollection *mongo.Collection

func init() {
	db := config.Connect()
	companyCollection = db.Collection("company")
}

func FindCompanyByEmail(email string) m.Company {
	result := m.Company{}
	err := companyCollection.FindOne(context.TODO(), bson.M{"companyemail": email}).Decode(&result)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return result
	}
	return result
}

func InsertOneCompany(company []byte) (*mongo.InsertOneResult, error) {
	resultData, err := companyCollection.InsertOne(config.Ctx, company)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		return nil, err
	}
	return resultData, nil
}
