package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lleonesouza/bus-travel/handlers"
	"github.com/lleonesouza/bus-travel/prisma/db"
	"github.com/lleonesouza/bus-travel/services"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/lleonesouza/bus-travel/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title bus-travel
// @version 1.0
// @description This is a simple Product (bus-travel) CRUD server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/lleonesouza
// @contact.email lleonesouza@live.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1324
// @BasePath /
func main() {
	e := echo.New()

	// Init DB
	client := db.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		panic(err)
	}

	// Initialize the bus travel service
	busTravelService := services.NewBusTravelService(client)

	// Initialize the bus travel handler with the service
	busTravelHandler := handlers.MakeHandlers(busTravelService)

	// Create Group and Register Routes
	busTravelGroup := e.Group("/bus-travels")
	busTravelHandler.BusTravelHandler.RegisterRoutes(busTravelGroup)

	// Init Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Listening Port
	e.Logger.Fatal(e.Start(":1324"))
}
