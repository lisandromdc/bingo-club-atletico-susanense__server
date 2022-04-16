package sellerRoutes

import (
	"github.com/gofiber/fiber/v2"
	sellerHandler "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/seller"
)

func SetupRoutes(router fiber.Router) {
	seller := router.Group("/seller")
	seller.Get("/", sellerHandler.GetSellers)
	seller.Get("/:sellerId", sellerHandler.GetOneSeller)
	seller.Post("/", sellerHandler.CreateSeller)
	seller.Put("/:sellerId", sellerHandler.UpdateSeller)
	seller.Delete("/:sellerId", sellerHandler.DeleteSeller)
}
