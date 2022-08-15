package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env.local")
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return
	} else if os.Getenv("ENV") == "prod" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return
	}

	err = godotenv.Load(".env.local.db")
	if err != nil || os.Getenv("ENV") == "prod" {
		err := godotenv.Load(".env.db")
		if err != nil {
			log.Fatal("Error loading .env.db file")
		}
		return
	} else if os.Getenv("ENV") == "prod" {
		err := godotenv.Load(".env.db")
		if err != nil {
			log.Fatal("Error loading .env.db file")
		}
		return
	}
}
