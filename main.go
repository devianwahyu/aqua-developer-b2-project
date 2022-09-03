package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")

	log.Println(dbHost)
}
