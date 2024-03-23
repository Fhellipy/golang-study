package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database
var Client *mongo.Client

func InitDB() {
	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		log.Fatal("DATABASE_URI is empty")
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("DATABASE_NAME is empty")
	}

	Database = client.Database(dbName)
	Client = client
}
