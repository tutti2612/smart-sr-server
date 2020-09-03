package main

import (
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/models"
)

func main() {
	db := database.Connection()
	db.AutoMigrate(&models.Student{})
}
