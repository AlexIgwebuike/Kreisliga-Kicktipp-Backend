package controller

import (
	"net/http"

	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/server"
	"github.com/AlexIgwebuike/Kreisliga-Kicktipp/service"
	"github.com/labstack/echo/v4"
)

var userApiRoute = server.ApiRoute + "/users"

func createUser(context echo.Context) error {

	var createUserRequest struct {
		Vorname  string
		Nachname string
		Email    string
		Passwort string
	}

	if err := context.Bind(&createUserRequest); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createUserResult, createUserError := service.CreateUser(createUserRequest.Vorname, createUserRequest.Nachname, createUserRequest.Email, createUserRequest.Passwort)

	if createUserError != nil {
		//log.Printf("Failed to create User %v", createUserError)

		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create User"})
	}

	return context.JSON(http.StatusCreated, map[string]interface{}{
		"userID": createUserResult.InsertedID,
	})

}

func SetupUserRoutes(echoServer *echo.Echo) {
	echoServer.POST(userApiRoute, createUser)
}
