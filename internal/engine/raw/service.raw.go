package raw

import (
	"codegen-service/internal/engine/models"
)

type ServiceMethod struct {
	Type        string
	Name        string
	ResponseDTO string
	DTOFields   []string
	Endpoint    *models.Endpoint
}

type ServiceRawData struct {
	ModelName   string
	ModelNameUC string
	Methods     []ServiceMethod
	General     models.General
	Fields      []models.Field // Все поля модели
}
