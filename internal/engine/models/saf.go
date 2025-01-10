package models

type Saf struct {
	General General `json:"general"`
	Models  []Model `json:"models"`
}
