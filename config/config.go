package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects go to postgres database
func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()

	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser, dbPass, dbHost, dbName, port := os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw", dbHost, dbUser, dbPass, dbName, port)

	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect postgres database")
	}

	return db
}

// DisconnectDB is stopping your connection to postgres database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to kill connection from database")
	}

	dbSQL.Close()
}
