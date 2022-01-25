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

func Connect() *mongo.Database {
	url := s.Join([]string{os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")}, ":")
	mongoURl := s.Join([]string{"mongodb://", url}, "")
	clientOptions := options.Client().ApplyURI(mongoURl)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Mongo Db Connected!")
	}
	db := client.Database(os.Getenv("MONGO_DATABASE_NAME"))
	return db
}
