package saleRoutes

import (
	"github.com/gofiber/fiber/v2"
	saleHandler "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/sale"
)

func SetupRoutes(router fiber.Router) {
	sale := router.Group("/sale")
	sale.Get("/", saleHandler.GetSales)
	sale.Get("/salesSummary", saleHandler.GetSalesSummary)
	sale.Put("/:saleId", saleHandler.UpdateSale)
	sale.Post("/", saleHandler.CreateSale)
	sale.Delete("/:saleId", saleHandler.DeleteSale)
}
