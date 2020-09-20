package main

import (
	"github.com/joho/godotenv"
	"github.com/tutti2612/smart-sr-server/internal/router"
	"log"
	"os"
)

func init() {
	envFile := ".env"
	if env := os.Getenv("GO_ENV"); env != "" {
		envFile += "." + env
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}
}

func main() {
	router.Run()
}
