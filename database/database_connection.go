package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientIstance *mongo.Client
var clientIstanceError error
var databaseClientSingleton sync.Once

func loadDatabaseUriFromEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	mongoDBUri := os.Getenv("MONGODB_URI")

	return mongoDBUri
}

func CreateDatabaseClient() (*mongo.Client, error) {
	databaseClientSingleton.Do(func() {
		databaseUri := loadDatabaseUriFromEnv()

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))

		if err != nil {
			clientIstanceError = err
			panic(clientIstanceError)
		}

		clientIstance = client

	})

	return clientIstance, clientIstanceError
}

func DisconnectMongoClient() error {
	if clientIstance != nil {
		clientIstance.Disconnect(context.TODO())
	}

	return nil
}
