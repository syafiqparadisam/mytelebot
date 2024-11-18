package entity

import "time"

type MessagePayload struct {
	// Id         *int       `json:"id"`
	Message string `json:"message"`
	ChatId  int64  `json:"chat_id"`
	// Created_at *time.Time `json:"created_at"`
}

type Message struct {
	Id         int       `json:"id"`
	Message    string    `json:"message"`
	ChatId     int64     `json:"chat_id"`
	Created_at time.Time `json:"created_at"`
}
