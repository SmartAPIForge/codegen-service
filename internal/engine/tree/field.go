package tree

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	IsUnique bool   `json:"isUnique"`
}
