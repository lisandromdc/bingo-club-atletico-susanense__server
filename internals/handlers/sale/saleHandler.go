package saleHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	dbCalls "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/hooks"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/gorm"
)

type saleModelGet struct {
	model.Sale
	// seller.DelegationId int
	// Delegation   model.Delegation
}

func preloadSale() *gorm.DB {
	m := []model.Sale{}
	query := database.DB
	query = query.Model(&m)
	query = query.Joins("LEFT JOIN sellers AS Seller ON sales.seller_id = Seller.id")
	query = query.Joins("LEFT JOIN delegations AS Delegation ON Seller.delegation_id = Delegation.id")
	query = query.Preload("Seller")
	query = query.Preload("Seller.Delegation")
	return query
}

func GetSales(c *fiber.Ctx) error {
	result := []saleModelGet{}
	params := struct {
		DelegationId int `json:"delegationId"`
		SellerId     int `json:"sellerId"`
		BingoNumber  int `json:"bingoNumber"`
		BingoId      int `json:"bingoId"`

		Search string `json:"search"`
	}{}
	return dbCalls.GetTableRow(c, &result, &params, func() *gorm.DB {
		s := "%" + params.Search + "%"
		query := preloadSale()
		query = query.Where("buyer_name ILIKE ?", s)
		query = query.Where(model.Sale{SellerId: params.SellerId})
		query = query.Where(model.Sale{BingoId: params.BingoId})
		query = query.Where(model.Sale{BingoNumber: params.BingoNumber})
		if params.DelegationId != 0 {
			query = query.Where("Seller.delegation_id = ?", params.DelegationId)
		}
		query = query.Order("created_at DESC")
		query = query.Find(&result)
		return query
	})
}

func GetSalesSummary(c *fiber.Ctx) error {
	result := []struct {
		DelegationId  int
		TotalQuantity int
		Delegation    model.Delegation
		SalePaymentId int
		IsPayed       bool
	}{}
	m := []model.Sale{}
	params := struct {
		BingoId int `json:"bingoId"`
	}{}
	return dbCalls.GetTableRow(c, &result, &params, func() *gorm.DB {
		query := database.DB
		query = query.Model(&m)
		query = query.Select(`
			Seller.delegation_id,
			COUNT(sales.id) AS total_quantity,
			SalePayment.id AS sale_payment_id,
			CASE WHEN SalePayment.id IS NULL THEN false ELSE true END as is_payed
		`)
		query = query.Joins("LEFT JOIN sellers AS Seller ON sales.seller_id = Seller.id")
		query = query.Joins("LEFT JOIN delegations AS Delegation ON Seller.delegation_id = Delegation.id")
		query = query.Joins("LEFT JOIN sale_payments AS SalePayment ON SalePayment.delegation_id = Delegation.id AND SalePayment.deleted_at IS NULL")
		query = query.Preload("Delegation")
		query = query.Where(model.Sale{BingoId: params.BingoId})
		query = query.Group(`Seller.delegation_id, Delegation.id, SalePayment.id`)
		query = query.Order("COUNT(sales.id) DESC")
		query = query.Find(&result)
		return query
	})
}

func CreateSale(c *fiber.Ctx) error {
	m := model.Sale{}
	result := saleModelGet{}
	return dbCalls.CreateTableRow(c, &m, func() *gorm.DB {
		query := preloadSale()
		query.First(result, m.ID)
		return query
	})
}

func DeleteSale(c *fiber.Ctx) error {
	m := model.Sale{}
	result := saleModelGet{}
	return dbCalls.DeleteTableRow(c, &m, "sales", "saleId", &result, preloadSale)
}

func UpdateSale(c *fiber.Ctx) error {
	m := model.Sale{}
	result := saleModelGet{}
	params := struct {
		BingoNumber  int    `json:"bingoNumber"`
		BuyerName    string `json:"buyerName"`
		BuyerAddress string `json:"buyerAddress"`
		BuyerPhone   string `json:"buyerPhone"`
	}{}
	return dbCalls.UpdateTableRow(c, &m, "sales", "saleId", &result, &params, preloadSale, func() {
		m.BingoNumber = params.BingoNumber
		m.BuyerName = params.BuyerName
		m.BuyerAddress = params.BuyerAddress
		m.BuyerPhone = params.BuyerPhone
	})
}
