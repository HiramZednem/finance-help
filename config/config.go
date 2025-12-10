package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading enviroment variables: ", err)
	}

	telegram_token := os.Getenv("telegram_token")
	if telegram_token == "" {
		log.Fatal("Bot Token not available")
		panic("Bot Token not available")
	}

	return &Config{
		TelegramToken: telegram_token,
	}
}
