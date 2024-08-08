package server

import (
	"log"
	"sync"

	"github.com/labstack/echo/v4"
)

var echoServerInstance *echo.Echo
var echoServerSingleton sync.Once

var ApiRoute string = "/api/v0"

func CreateEchoServer() *echo.Echo {
	echoServerSingleton.Do(func() {
		echoServerInstance = echo.New()

	})

	return echoServerInstance
}

func StartEchoServer(echoServer *echo.Echo) {
	log.Println("Starting Server on Port :8080")
	if err := echoServer.Start(":8080"); err != nil {
		log.Fatalf("Error starting server %v", err)
	}
}
