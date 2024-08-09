package main

import (
	"log"

	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/controller"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/database"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/server"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var databaseClient *mongo.Client
var databaseClientError error

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}

func run() error {
	if err := LoadEnvFile(); err != nil {
		return err
	}

	if err := initDatabase(); err != nil {
		return err
	}

	initServer()

	return nil
}

func LoadEnvFile() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env File")
		return err
	}

	return nil
}

func initDatabase() error {
	databaseClient, databaseClientError = database.CreateDatabaseClient()

	if databaseClientError != nil {
		log.Printf("Failed to Connect to Database %v", databaseClientError)
	}

	log.Println("Connected Successfully to Database")

	service.SetupDatabaseUserService(databaseClient)

	return nil
}

func initServer() {
	echoServer := server.CreateEchoServer()

	controller.SetupUserRoutes(echoServer)

	server.StartEchoServer(echoServer)
}
