package bingoHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	dbCalls "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/hooks"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/gorm"
)

type bingoModelGet = struct {
	model.Bingo
	SalesQuantity int
}

func preloadBingo() *gorm.DB {
	m := []model.Bingo{}
	query := database.DB
	query = query.Model(&m)
	return query
}

func GetBingos(c *fiber.Ctx) error {
	m := []model.Bingo{}
	result := []bingoModelGet{}
	params := struct {
		DelegationId int `json:"delegationId"`
	}{}
	return dbCalls.GetTableRow(c, &result, &params, func() *gorm.DB {
		query := database.DB
		query = query.Model(&m)
		query = query.Select("bingos.*, COUNT(Sale.id) AS sales_quantity")
		query = query.Joins("LEFT JOIN sales AS Sale ON Sale.bingo_id = bingos.id")
		query = query.Joins("LEFT JOIN sellers AS Seller ON Sale.seller_id = Seller.id")
		if params.DelegationId != 0 {
			query = query.Where("Seller.delegation_id = ?", params.DelegationId)
		}
		query = query.Group("Sale.bingo_id, bingos.id")
		query = query.Order("created_at DESC")
		query = query.Find(&result)
		return query
	})
}

func CreateBingo(c *fiber.Ctx) error {
	m := model.Bingo{}
	result := bingoModelGet{}
	return dbCalls.CreateTableRow(c, &m, func() *gorm.DB {
		query := preloadBingo()
		query.First(result, m.ID)
		return query
	})
}

func DeleteBingo(c *fiber.Ctx) error {
	m := model.Bingo{}
	result := bingoModelGet{}
	return dbCalls.DeleteTableRow(c, &m, "bingos", "bingoId", &result, preloadBingo)
}
