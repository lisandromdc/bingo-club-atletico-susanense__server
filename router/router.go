package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	bingoRoutes "github.com/lisandromdc/bingo-club-atletico-susanense/router/bingo"
	delegationRoutes "github.com/lisandromdc/bingo-club-atletico-susanense/router/delegation"
	saleRoutes "github.com/lisandromdc/bingo-club-atletico-susanense/router/sale"
	salePaymentRoutes "github.com/lisandromdc/bingo-club-atletico-susanense/router/sale-payment"
	sellerRoutes "github.com/lisandromdc/bingo-club-atletico-susanense/router/seller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	bingoRoutes.SetupRoutes(api)
	delegationRoutes.SetupRoutes(api)
	saleRoutes.SetupRoutes(api)
	salePaymentRoutes.SetupRoutes(api)
	sellerRoutes.SetupRoutes(api)
}
