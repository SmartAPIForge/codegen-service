package models

type Field struct {
	Primary  bool   `json:"primary"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	IsUnique bool   `json:"isUnique"`
}
