package bingoRoutes

import (
	"github.com/gofiber/fiber/v2"
	bingoHandler "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/bingo"
)

func SetupRoutes(router fiber.Router) {
	bingo := router.Group("/bingo")
	bingo.Get("/", bingoHandler.GetBingos)
	bingo.Post("/", bingoHandler.CreateBingo)
	bingo.Delete("/:bingoId", bingoHandler.DeleteBingo)
}
