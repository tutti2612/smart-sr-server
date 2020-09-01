package main

import (
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/model"
)

func main() {
	db := database.Connection()
	db.AutoMigrate(&model.Student{})
}
