package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateController(model *models.Model) {
	templateName := "controller.tmpl"
	pathToDir, err := g.CreateDir(fmt.Sprintf("%s/%s", g.projectRoot, model.Name))
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/%s.controller.go", pathToDir, model.Name))
	defer outputFile.Close()

	rawData := g.fetchControllerRawData(model)
	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) fetchControllerRawData(model *models.Model) *raw.ControllerRawData {
	var routes []raw.Route
	for _, endpoint := range model.Endpoints {
		route := raw.Route{
			Method: endpoint.Type,
		}
		routes = append(routes, route)
	}
	rawData := raw.ControllerRawData{
		ModelName:   model.Name,
		ModelNameUC: g.ToUC(model.Name),
		Routes:      routes,
	}

	return &rawData
}
