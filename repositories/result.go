package repositories

import "time"

type Result struct {
	Id int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username string `json:"username"`
}

