package raw

type FieldRaw struct {
	Name string
	Type string
}

type ModelRawData struct {
	ModelName   string
	ModelNameUC string
	Fields      []FieldRaw
}
