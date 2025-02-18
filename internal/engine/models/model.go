package models

type Model struct {
	Name      string     `json:"name"`
	Fields    []Field    `json:"fields"`
	Endpoints []Endpoint `json:"endpoints"`
}
