package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	TelegramApiEndpoint string
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

	telegram_api_endpoint := os.Getenv("telegram_api_endpoint")
	if telegram_api_endpoint == "" {
		log.Fatal("Telegram API Endpoint not available")
		panic("Telegram API Endpoint not available")
	}


	return &Config{
		TelegramToken: telegram_token,
		TelegramApiEndpoint: telegram_api_endpoint,
	}
}
