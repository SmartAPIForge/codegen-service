package tree

type Method struct {
	Type            string   `json:"type"`
	UniqueParam     string   `json:"uniqueParam"`
	ResponseFields  []string `json:"responseFields"`
	PrivateEndpoint bool     `json:"privateEndpoint"`
}
