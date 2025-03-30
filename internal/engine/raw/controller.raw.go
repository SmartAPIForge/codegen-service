package raw

type Route struct {
	Method      string
	HandlerName string
	IsRegistered bool
}

type ControllerRawData struct {
	ModelName   string
	ModelNameUC string
	Routes      []Route
}
