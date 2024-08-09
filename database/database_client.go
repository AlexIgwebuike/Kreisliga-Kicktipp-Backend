package database

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientIstance *mongo.Client
var clientIstanceError error
var databaseClientSingleton sync.Once

func loadDatabaseUriFromEnv() string {
	mongoDBUri := os.Getenv("MONGODB_URI")

	return mongoDBUri
}

func CreateDatabaseClient() (*mongo.Client, error) {
	databaseClientSingleton.Do(func() {
		databaseUri := loadDatabaseUriFromEnv()

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(databaseUri))

		if err != nil {
			clientIstanceError = err
			log.Printf("Error %v", clientIstanceError)

			return
		}

		clientIstance = client

	})

	return clientIstance, clientIstanceError
}

func CreatUniqueIndex(collection *mongo.Collection, field string) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{field: 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func DisconnectDatabaseClient() error {
	if clientIstance != nil {
		clientIstance.Disconnect(context.TODO())
	}

	return nil
}
