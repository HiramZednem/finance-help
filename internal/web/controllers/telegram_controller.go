package controllers

import (
	"encoding/json"
	"finance-help/internal/web/dtos"
	"io"
	"log"
	"net/http"
)

type TelegramControllerImpl struct {
	// TODO: add service
}

type TelegramControllerInterface interface {
	HandleEvent(w http.ResponseWriter, r *http.Request)
}

func NewTelegramController() TelegramControllerInterface {
	return &TelegramControllerImpl{}
}

func (t *TelegramControllerImpl) HandleEvent(w http.ResponseWriter, r *http.Request) {
	// TODO: add origin verification...

	


	telegramBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Err reading body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	telegramEvent, err := dtos.UnmarshalTelegramEvent(telegramBody)
	if err != nil {
		log.Println("Err unmarshaling tg event ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Message from: %s\nContent: %s", telegramEvent.Message.From.Username, telegramEvent.Message.Text)	

	response := map[string]interface{}{
		"method": "sendMessage",
		"chat_id": telegramEvent.Message.Chat.ID,
		"text": "Testing",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		log.Println("Err marshaling body: ", err)
	}
	w.Write(responseBytes)

		
	
}