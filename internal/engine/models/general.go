package models

type General struct {
	Id       string
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Port     int    `json:"port"`
	AutoAuth bool   `json:"autoAuth"`
}
