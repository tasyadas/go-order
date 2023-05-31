package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "tasyadas"
	dbPort = "5432"
	dbName = "db_inventory_sql"
	db     *gorm.DB
	err    error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.AutoMigrate(&models.Order{}, &models.Item{})

}

func GetDB() *gorm.DB {
	return db
}
