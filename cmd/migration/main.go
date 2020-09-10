package main

import (
	"github.com/joho/godotenv"
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/models"
	"log"
	"os"
)

func main() {
	envFile := ".env"
	if env := os.Getenv("GO_ENV"); env != "" {
		envFile += "." + env
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}

	db := database.Connection()
	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Fatal(err)
	}
}
