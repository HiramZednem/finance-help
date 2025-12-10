package main

import (
	"finance-help/config"
	"finance-help/internal/handlers"
	"log"
	"os"
	"os/signal"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("Config Loaded")

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal("Err creating the bot: ", err)
	}
	log.Println("Bot created")

	tgHandler := handlers.NewTelegramHandler(bot)

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)


	// TODO: improve this
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	log.Println("Setting Update Handler")

	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				tgHandler.HandleUpdate(update)
			}

		case <-sig:
			log.Println("Stopping bot...")

			bot.StopReceivingUpdates()

			log.Println("Bot stopped cleanly.")
			return
		}
	}
}
