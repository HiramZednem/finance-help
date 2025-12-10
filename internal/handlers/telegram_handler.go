package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHandlerImpl struct {
	bot *tgbotapi.BotAPI
}

type TelegramHandlerInterface interface {
	HandleUpdate(update tgbotapi.Update)
}

func NewTelegramHandler(bot *tgbotapi.BotAPI) TelegramHandlerInterface {
	return &TelegramHandlerImpl{
		bot: bot,
	}
}

func (t *TelegramHandlerImpl) HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	log.Println("Handling update")
 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Message recieved")
	t.bot.Send(msg)
}