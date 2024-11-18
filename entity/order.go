package entity

type Order struct {
	ChatId int64 `json:"chat_id"`
	Price int64 `json:"price"`
	ServicesType string `json:"services_type"`
	ServicesId int `json:"services_id"`
}