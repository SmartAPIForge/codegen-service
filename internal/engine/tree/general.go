package tree

type General struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Owner    string `json:"owner"`
	Version  string `json:"version"`
	Port     int    `json:"port"`
	AutoAuth bool   `json:"autoAuth"`
}
