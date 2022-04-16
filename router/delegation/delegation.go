package delegationRoutes

import (
	"github.com/gofiber/fiber/v2"
	delegationHandler "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/delegation"
)

func SetupRoutes(router fiber.Router) {
	delegation := router.Group("/delegation")
	delegation.Get("/", delegationHandler.GetDelegations)
	delegation.Post("/", delegationHandler.CreateDelegation)
	delegation.Put("/:delegationId", delegationHandler.UpdateDelegation)
	delegation.Delete("/:delegationId", delegationHandler.DeleteDelegation)
}
