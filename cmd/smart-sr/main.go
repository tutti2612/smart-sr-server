package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tutti2612/smart-sr-server/internal/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}

	router.Run()
}
