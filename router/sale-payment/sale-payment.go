package salePaymentRoutes

import (
	"github.com/gofiber/fiber/v2"
	salePaymentHandler "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/sale-payment"
)

func SetupRoutes(router fiber.Router) {
	salePayment := router.Group("/salePayment")
	salePayment.Post("/", salePaymentHandler.CreateSalePayment)
	salePayment.Delete("/:salePaymentId", salePaymentHandler.DeleteSalePayment)
}
