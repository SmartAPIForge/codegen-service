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

	// Карта для отслеживания эндпоинтов, для которых уже созданы DTO
	existingDTOs := make(map[string]bool)

	// Сначала создаем DTO для существующих эндпоинтов
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
		
		existingDTOs[endpoint.Type] = true
	}
	
	// Добавляем DTO для стандартных типов запросов, если они отсутствуют
	standardTypes := []string{"GET", "POST", "DELETE"}
	for _, dtoType := range standardTypes {
		if !existingDTOs[dtoType] {
			// Создаем минимальный эндпоинт с полем id
			dummyEndpoint := models.Endpoint{
				Type:        dtoType,
				ResponseDTO: []string{"id"},
			}
			
			outputFile, _ := g.CreateFile(fmt.Sprintf(
				"%s/%s%s.dto.go",
				pathToDir,
				g.ToLower(dtoType),
				g.ToUC(model.Name),
			))
			
			rawData := g.fetchDTORawData(
				model,
				&dummyEndpoint,
			)
			err := g.templates.ExecuteTemplate(outputFile, templateName, rawData)
			if err != nil {
				panic(err)
			}
			outputFile.Close()
		}
	}
}

func (g *Generator) fetchDTORawData(
	model *models.Model,
	endpoint *models.Endpoint,
) *raw.DTORawData {
	var requiredFields []raw.FieldRaw
	// Всегда добавляем поле id в DTO
	idField := raw.FieldRaw{
		Name: "Id",
		Type: "int",
	}
	requiredFields = append(requiredFields, idField)
	
	// Добавляем остальные поля из responseDTO
	for _, field := range model.Fields {
		// Пропускаем id, так как мы уже добавили его
		if field.Name == "id" {
			continue
		}
		
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
