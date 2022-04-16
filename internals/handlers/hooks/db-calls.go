package dbCalls

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lisandromdc/bingo-club-atletico-susanense/database"
	"gorm.io/gorm"
)

// INNER
func emptyPreload() *gorm.DB {
	return database.DB
}

// EXTERNAL
func GetOneRow(
	c *fiber.Ctx,
	key string,
	resultModel interface{},
	preloadFunc func() *gorm.DB,
) error {
	query := preloadFunc()
	query.First(resultModel, c.Params(key))

	if query.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get", "data": query.Error})
	}

	return c.JSON(fiber.Map{"status": "success", "data": resultModel})
}
func GetTableRow(
	c *fiber.Ctx,
	tableModel interface{},
	where interface{},
	cb func() *gorm.DB,
) error {
	// Store the body containing the updated data and return error if encountered
	err := c.QueryParser(where)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	query := cb()

	return c.JSON(fiber.Map{"status": "success", "data": tableModel, "count": query.RowsAffected})
}

func CreateTableRow(
	c *fiber.Ctx,
	tableModel interface{},
	cb func() *gorm.DB,
) error {
	err := c.BodyParser(tableModel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	res := database.DB.Create(tableModel)
	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create", "data": err})
	}

	result := cb()

	return c.JSON(fiber.Map{"status": "success", "message": "Created", "data": result})
}

func UpdateTableRow(
	c *fiber.Ctx,
	tableModel interface{},
	tableName string,
	key string,
	resultModel interface{},
	params interface{},
	preloadFunc func() *gorm.DB,
	cb func(),
) error {
	findById(tableModel, tableName, c.Params(key), emptyPreload)

	// if tableModel.ID == 0 {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found", "data": nil})
	// }

	// Store the body containing the updated data and return error if encountered
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the seller
	// for _, fieldName := range paramsFields {
	// 	SetField(&resultModel, fieldName, getAttr(&params, fieldName))
	// }
	cb()

	// Save the Changes
	database.DB.Save(tableModel)

	// get new model
	findById(resultModel, tableName, c.Params(key), preloadFunc)

	// Return the updated seller
	return c.JSON(fiber.Map{"status": "success", "message": "Found", "data": resultModel})
}

func findById(
	resultModel interface{},
	tableName string,
	keyValue string,
	preloadFunc func() *gorm.DB,
) *gorm.DB {
	query := preloadFunc()
	query = query.Find(resultModel, tableName+".id = ?", keyValue)
	return query
}

func DeleteTableRow(
	c *fiber.Ctx,
	tableModel interface{},
	tableName string,
	key string,
	resultModel interface{},
	preloadFunc func() *gorm.DB,
) error {

	findById(resultModel, tableName, c.Params(key), preloadFunc)

	// if resultModel.ID == 0 {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Not found", "data": nil})
	// }

	// Delete the note and return error if encountered
	err := database.DB.Delete(tableModel, "id = ?", c.Params(key)).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted"})
}
