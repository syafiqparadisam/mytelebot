package entity

type Os struct {
	Id          int     `json:"id"`
	Distro      string  `json:"distro"`
	Description *string `json:"description"`
	Level       string  `json:"level"`
	Price       int64   `json:"price"`
}
