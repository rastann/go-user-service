package database

import (
	"fmt"
	"log"

	"tutorial/user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB) {
	config := config.GetConfig()
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.Username, config.Password, config.DBName, config.Port, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database")
		return
	}
	return db
}
