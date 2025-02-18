package models

type General struct {
	Id       string
	Port     int
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	AutoAuth bool   `json:"autoAuth"`
}
