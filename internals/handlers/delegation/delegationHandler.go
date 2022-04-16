package delegationHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	dbCalls "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/hooks"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/gorm"
)

type delegationModelGet = struct {
	model.Delegation
	SellersQuantity int
}

func preloadDelegation() *gorm.DB {
	m := []model.Delegation{}
	query := database.DB
	query = query.Model(&m)
	query = query.Select("delegations.*, COUNT(Seller.id) AS sellers_quantity")
	query = query.Joins("LEFT JOIN sellers AS Seller ON Seller.delegation_id = delegations.id")
	query = query.Group("Seller.delegation_id, delegations.id")
	return query
}

func GetDelegations(c *fiber.Ctx) error {
	result := []delegationModelGet{}
	params := struct {
		Search string `json:"search"`
	}{}
	return dbCalls.GetTableRow(c, &result, &params, func() *gorm.DB {
		query := preloadDelegation()
		query = query.Where("delegations.name ILIKE ?", "%"+params.Search+"%")
		query = query.Order("delegations.name ASC")
		query = query.Find(&result)
		return query
	})
}

func CreateDelegation(c *fiber.Ctx) error {
	m := model.Delegation{}
	result := delegationModelGet{}
	return dbCalls.CreateTableRow(c, &m, func() *gorm.DB {
		query := preloadDelegation()
		query.First(result, m.ID)
		return query
	})
}

func UpdateDelegation(c *fiber.Ctx) error {
	m := model.Delegation{}
	result := delegationModelGet{}
	params := struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Notes string `json:"notes"`
	}{}
	return dbCalls.UpdateTableRow(c, &m, "delegations", "delegationId", &result, &params, preloadDelegation, func() {
		m.Name = params.Name
		m.Icon = params.Icon
		m.Notes = params.Notes
	})
}

func DeleteDelegation(c *fiber.Ctx) error {
	m := model.Delegation{}
	result := delegationModelGet{}
	return dbCalls.DeleteTableRow(c, &m, "delegations", "delegationId", &result, preloadDelegation)
}
