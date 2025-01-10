package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateMain(saf *models.Saf) {
	templateName := "main.tmpl"
	pathToDir, err := g.CreateDir(g.projectRoot)
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/main.go", pathToDir))
	defer outputFile.Close()

	rawData := g.fetchMainRawData(saf)
	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) fetchMainRawData(saf *models.Saf) *raw.MainRawData {
	var extendedModels []raw.ExtendedModel
	for _, model := range saf.Models {
		extendedModel := raw.ExtendedModel{
			Model:  model,
			NameUC: g.ToUC(model.Name),
		}
		extendedModels = append(extendedModels, extendedModel)
	}
	rawData := raw.MainRawData{
		Models:  extendedModels,
		General: saf.General,
	}

	return &rawData
}
