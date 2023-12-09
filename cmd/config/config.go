package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BOT_TOKEN string

const PREFIX = "!bot"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BOT_TOKEN = os.Getenv("BOT_TOKEN")
}
