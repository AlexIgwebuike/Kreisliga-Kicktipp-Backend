package main

import (
	"log"

	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/controller"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/database"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/server"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var databaseClient *mongo.Client
var databaseClientError error

func main() {

	//Setup Database
	databaseClient, databaseClientError = database.CreateDatabaseClient()

	if databaseClientError != nil {
		log.Fatalf("Failed to Connect to Database %v", databaseClientError)
	}

	log.Println("Connected Successfully to Database")

	service.SetupDatabaseUserService(databaseClient)

	//Start Echo Server
	echoServer := server.CreateEchoServer()
	controller.SetupUserRoutes(echoServer)
	server.StartEchoServer(echoServer)

}
