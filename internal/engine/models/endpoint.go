package models

type Endpoint struct {
	Type        string   `json:"type"`
	Query       []string `json:"query"`
	ResponseDTO []string `json:"responseDTO"`
	Private     bool     `json:"privateEndpoint"`
	IsRegistered bool    `json:"isRegistered"`
}
