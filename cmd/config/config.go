package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const PREFIX = "!bot"

type Config struct {
	BOT_TOKEN string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bt := os.Getenv("BOT_TOKEN")

	return Config{
		BOT_TOKEN: bt,
	}

}
