package salePaymentHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	dbCalls "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/hooks"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/gorm"
)

type salePaymentModelGet struct {
	model.SalePayment
}

func preloadSalePayment() *gorm.DB {
	m := []model.SalePayment{}
	query := database.DB
	query = query.Model(&m)
	return query
}

func CreateSalePayment(c *fiber.Ctx) error {
	m := model.SalePayment{}
	result := salePaymentModelGet{}
	return dbCalls.CreateTableRow(c, &m, func() *gorm.DB {
		query := preloadSalePayment()
		query.First(result, m.ID)
		return query
	})
}

func DeleteSalePayment(c *fiber.Ctx) error {
	m := model.SalePayment{}
	result := salePaymentModelGet{}
	return dbCalls.DeleteTableRow(c, &m, "SalePayments", "salePaymentId", &result, preloadSalePayment)
}
