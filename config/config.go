package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB connects go to postgres database
func ConnectDB() *gorm.DB {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUser, dbPass, dbHost, dbName, dbPort := os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw", dbHost, dbUser, dbPass, dbName, dbPort)

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
