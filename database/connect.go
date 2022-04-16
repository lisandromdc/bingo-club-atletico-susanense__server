package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lisandromdc/bingo-club-atletico-susanense/config"
	"github.com/lisandromdc/bingo-club-atletico-susanense/internals/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	fmt.Println("port", port)
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn))

	fmt.Println("DB", DB)
	fmt.Println("err", err)

	// Connect to the DB and initialize the DB variable

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	DB.AutoMigrate(&model.Bingo{})
	DB.AutoMigrate(&model.Delegation{})
	DB.AutoMigrate(&model.Sale{})
	DB.AutoMigrate(&model.SalePayment{})
	DB.AutoMigrate(&model.Seller{})
	fmt.Println("Database Migrated")
}
