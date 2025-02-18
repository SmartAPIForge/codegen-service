package raw

type Route struct {
	Method      string
	HandlerName string
}

type ControllerRawData struct {
	ModelName   string
	ModelNameUC string
	Routes      []Route
}
