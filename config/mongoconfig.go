package config

import (
	"context"
	"log"
	"os"
	s "strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var url string

func Connect() *mongo.Database {
	if os.Getenv("ENVIRONMENT") == "testing" {
		url = s.Join([]string{os.Getenv("MONGO_TEST_HOST"), os.Getenv("MONGO_TEST_PORT")}, ":")
		if os.Getenv("MONGO_TEST_USERNAME") != "" && os.Getenv("MONGO_TEST_PASSWORD") != "" {
			url = s.Join([]string{s.Join([]string{os.Getenv("MONGO_TEST_USERNAME"), os.Getenv("MONGO_TEST_PASSWORD")}, ":"), url}, "@")
		}
	} else {
		url = s.Join([]string{os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")}, ":")
		if os.Getenv("MONGO_USERNAME") != "" && os.Getenv("MONGO_PASSWORD") != "" {
			url = s.Join([]string{s.Join([]string{os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD")}, ":"), url}, "@")
		}
	}
	mongoURl := s.Join([]string{"mongodb://", url}, "")
	clientOptions := options.Client().ApplyURI(mongoURl)
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Mongo Db Connected!")
	}
	if os.Getenv("ENVIRONMENT") == "testing" {
		db := client.Database(os.Getenv("MONGO_TEST_DATABASE_NAME"))
		return db
	} else {
		db := client.Database(os.Getenv("MONGO_DATABASE_NAME"))
		return db
	}
}
