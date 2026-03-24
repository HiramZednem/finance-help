package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	TelegramApiEndpoint string
	ENV string
	PORT string
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

	env := os.Getenv("env")
	if env == "" {
		log.Println("ENV not set, defaulting to 'dev'")
		env = "dev"
	}

	port := os.Getenv("port")
	if port == "" {
		log.Println("PORT not set, defaulting to '8080")
	}


	return &Config{
		TelegramToken: telegram_token,
		TelegramApiEndpoint: telegram_api_endpoint,
		ENV: env,
		PORT: port,
	}
}
