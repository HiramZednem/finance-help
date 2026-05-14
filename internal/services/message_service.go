package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MessageServiceInterface interface {
	SendMessage(message string) 
	ProcessMessage(message string) error
}

type BaseMessageService struct{}

func (b *BaseMessageService) ProcessMessage(message string) error {
	if message == "" {
		return errors.New("Message cannot be empty")
	}
	
	message = strings.Trim(message, " ")
	messages := strings.Split(message, " ")

	if len(messages) < 1 {
		return errors.New("Message should contain at least 2 words")
	}

	money, err := strconv.ParseFloat(messages[0], 32)
	if err != nil {
		return errors.New("First word should be a number")
	}

	concept := messages[1]

	fmt.Printf("Want to reduce: %f, with the concept: %s", money, concept)


	return nil
}
