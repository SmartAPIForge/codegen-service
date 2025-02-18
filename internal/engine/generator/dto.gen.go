package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
	"slices"
)

func (g *Generator) GenerateDTOs(model *models.Model) {
	templateName := "dto.tmpl"
	pathToDir, _ := g.CreateDir(fmt.Sprintf("%s/%s/dto", g.projectRoot, model.Name))

	for _, endpoint := range model.Endpoints {
		outputFile, _ := g.CreateFile(fmt.Sprintf(
			"%s/%s%s.dto.go",
			pathToDir,
			g.ToLower(endpoint.Type),
			g.ToUC(model.Name),
		))

		rawData := g.fetchDTORawData(
			model,
			&endpoint,
		)
		err := g.templates.ExecuteTemplate(outputFile, templateName, rawData)
		if err != nil {
			panic(err)
		}
		outputFile.Close()
	}
}

func (g *Generator) fetchDTORawData(
	model *models.Model,
	endpoint *models.Endpoint,
) *raw.DTORawData {
	var requiredFields []raw.FieldRaw
	for _, field := range model.Fields {
		if slices.Contains(endpoint.ResponseDTO, field.Name) {
			dtoField := raw.FieldRaw{
				Name: g.ToUC(field.Name),
				Type: field.Type,
			}
			requiredFields = append(requiredFields, dtoField)
		}
	}

	rawData := raw.DTORawData{
		ModelName: model.Name,
		DTOName:   fmt.Sprintf("%s%sDTO", g.ToUC(endpoint.Type), g.ToUC(model.Name)),
		Fields:    requiredFields,
	}

	return &rawData
}
