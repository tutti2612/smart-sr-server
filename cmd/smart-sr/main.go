package main

import (
	"github.com/joho/godotenv"
	"github.com/tutti2612/smart-sr-server/internal/router"
	"log"
	"os"
)

func main() {
	env := os.Getenv("GO_ENV")
	envFile := ".env"
	if env != "" {
		envFile += "." + env
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}

	router.Run()
}
