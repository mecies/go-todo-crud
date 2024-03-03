package main

import (
	"go-30/todo/config"
	"go-30/todo/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	routes.Routes()
}
