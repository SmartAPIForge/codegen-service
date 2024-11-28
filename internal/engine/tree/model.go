package tree

type Model struct {
	Name    string   `json:"name"`
	Fields  []Field  `json:"fields"`
	Methods []Method `json:"methods"`
}
