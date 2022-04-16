package sellerHandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	dbCalls "github.com/lisandromdc/bingo-club-atletico-susanense/internals/handlers/hooks"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/gorm"
)

type sellerModelGet struct {
	model.Seller
}

func preloadSeller() *gorm.DB {
	m := []model.Seller{}
	fmt.Println("database.DB", database.DB)
	query := database.DB
	query = query.Preload("Delegation")
	query = query.Model(&m)
	return query
}

func GetSellers(c *fiber.Ctx) error {
	result := []sellerModelGet{}
	params := struct {
		DelegationId      int    `json:"delegationId"`
		Email             string `json:"email"`
		Password          string `json:"password"`
		IncludeDeleted    bool   `json:"includeDeleted"`
		IsDelegationAdmin bool   `json:"isDelegationAdmin"`

		Search string `json:"search"`
	}{}
	// where := map[string]interface{}{"name": "Lisandro"}
	return dbCalls.GetTableRow(c, &result, &params, func() *gorm.DB {
		// database.DB.Where(&where).Find(&result)
		s := "%" + params.Search + "%"
		query := preloadSeller()
		query = query.Where("(CONCAT(name, ' ', surname) ILIKE ? OR email ILIKE ?)", s, s)
		query = query.Where(model.Seller{Email: params.Email})
		query = query.Where(model.Seller{Password: params.Password})
		query = query.Where(model.Seller{DelegationId: params.DelegationId})
		query = query.Where(model.Seller{IsDelegationAdmin: params.IsDelegationAdmin})
		if params.IncludeDeleted {
			query = query.Unscoped()
		}
		query = query.Order("CONCAT(surname, ' ', name) ASC")
		query = query.Find(&result)
		return query
	})
}

func GetOneSeller(c *fiber.Ctx) error {
	result := sellerModelGet{}
	return dbCalls.GetOneRow(c, "sellerId", &result, preloadSeller)
}

func CreateSeller(c *fiber.Ctx) error {
	fmt.Println("Creo")
	m := model.Seller{}
	// result := sellerModelGet{}
	return dbCalls.CreateTableRow(c, &m, func() *gorm.DB {
		query := preloadSeller()
		fmt.Println(m.ID)
		// query.First(result, m.ID)
		return query
	})
}

func UpdateSeller(c *fiber.Ctx) error {
	fmt.Println("Edito")
	m := model.Seller{}
	result := sellerModelGet{}
	params := struct {
		DelegationId      int    `json:"delegationId"`
		Name              string `json:"name"`
		Surname           string `json:"surname"`
		Email             string `json:"email"`
		Password          string `json:"password"`
		IsDelegationAdmin bool   `json:"isDelegationAdmin"`
	}{}
	return dbCalls.UpdateTableRow(c, &m, "sellers", "sellerId", &result, &params, preloadSeller, func() {
		m.DelegationId = params.DelegationId
		m.Name = params.Name
		m.Surname = params.Surname
		m.Email = params.Email
		m.Password = params.Password
		m.IsDelegationAdmin = params.IsDelegationAdmin
	})
}

func DeleteSeller(c *fiber.Ctx) error {
	m := model.Seller{}
	result := sellerModelGet{}
	return dbCalls.DeleteTableRow(c, &m, "sellers", "sellerId", &result, preloadSeller)
}
