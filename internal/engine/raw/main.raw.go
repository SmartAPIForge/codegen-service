package raw

import "codegen-service/internal/engine/models"

type ExtendedModel struct {
	models.Model
	NameUC string
}

type MainRawData struct {
	Models  []ExtendedModel
	General models.General
}
