package entity

type User struct {
	Username string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	ChatId int64 `json:"chat_id"`
}
