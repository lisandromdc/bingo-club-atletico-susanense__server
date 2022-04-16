package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	"github.com/lisandromdc/bingo-club-atletico-susanense/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()
	app.Use(cors.New())

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3030")
}
