package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateModel(model *models.Model) {
	templateName := "model.tmpl"
	pathToDir, _ := g.CreateDir(fmt.Sprintf("%s/%s", g.projectRoot, model.Name))
	outputFile, _ := g.CreateFile(fmt.Sprintf(
		"%s/%s.model.go",
		pathToDir,
		model.Name,
	))
	defer outputFile.Close()

	rawData := g.fetchModelRawData(model)
	err := g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) fetchModelRawData(
	model *models.Model,
) *raw.ModelRawData {
	var rawFields []raw.FieldRaw
	for _, field := range model.Fields {
		rawField := raw.FieldRaw{
			Name: g.ToUC(field.Name),
			Type: field.Type,
		}
		rawFields = append(rawFields, rawField)
	}

	rawData := raw.ModelRawData{
		ModelName:   model.Name,
		ModelNameUC: g.ToUC(model.Name),
		Fields:      rawFields,
	}

	return &rawData
}
