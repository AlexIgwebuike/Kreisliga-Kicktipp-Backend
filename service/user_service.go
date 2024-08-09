package service

import (
	"context"
	"fmt"
	"log"

	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/database"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/model"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var databaseClient *mongo.Client

func SetupDatabaseUserService(client *mongo.Client) {
	databaseClient = client
}

func CreateUser(vorname, nachname, email, password string) (*mongo.InsertOneResult, error) {
	userCollection := databaseClient.Database("Kicktipp").Collection("users")
	userPasswordCollection := databaseClient.Database("Kicktipp").Collection("user_passwords")
	hashedPassword, hashedPasswordError := util.HashPassword(password)

	if hashedPasswordError != nil {
		log.Fatalf("Failed to Hash Password: %v", hashedPasswordError)
	}

	uniqueIndexError := database.CreatUniqueIndex(userCollection, "email")

	if uniqueIndexError != nil {
		log.Fatalf("Failed to create Unique Index for the UserCollection: %v", uniqueIndexError)
	}

	user := model.User{
		ID:       primitive.NewObjectID(),
		Vorname:  vorname,
		Nachname: nachname,
		Email:    email,
	}

	result, userCollectionError := userCollection.InsertOne(context.TODO(), user)

	if userCollectionError != nil {
		if mongo.IsDuplicateKeyError(userCollectionError) {
			errorMessage := "email already exists"
			log.Println(errorMessage)
			return nil, fmt.Errorf(errorMessage)
		}
		log.Printf("Failed to access the User Collection %v", userCollectionError)
		return nil, userCollectionError
	}

	UserPassword := model.UserPassword{
		ID:             primitive.NewObjectID(),
		UserID:         user.ID,
		HashedPassword: hashedPassword,
	}

	_, userPasswordCollectionError := userPasswordCollection.InsertOne(context.TODO(), UserPassword)

	if userPasswordCollectionError != nil {
		log.Printf("Failed to acccess the User-Password Collection %v", userPasswordCollectionError)
		return nil, userPasswordCollectionError
	}

	log.Println("New User has been created!")
	return result, nil
}
