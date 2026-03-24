// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    telegramEvent, err := UnmarshalTelegramEvent(bytes)
//    bytes, err = telegramEvent.Marshal()
package dtos

import "encoding/json"

func UnmarshalTelegramEvent(data []byte) (TelegramRequest, error) {
	var r TelegramRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TelegramRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TelegramRequest struct {
	UpdateID int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	From      From   `json:"from"`
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	Text      string `json:"text"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type From struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
