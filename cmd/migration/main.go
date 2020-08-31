package main

import (
	"github.com/tutti2612/smart-sr-server/internal/db"
	"github.com/tutti2612/smart-sr-server/internal/model"
)

func main() {
	db := db.Connection()
	db.AutoMigrate(&model.Student{})
}