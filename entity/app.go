package entity

import "time"

type AppPayload struct {
	Description string `json:"description"`
}

type App struct {
	Id          int       `json:"id"`
	Deadline    time.Time `json:"deadline"`
	Description string    `json:"description"`
	Tech        string    `json:"tech"`
}

type UpdateTech struct {
	Id   int    `json:"id"`
	Tech string `json:"tech"`
}
