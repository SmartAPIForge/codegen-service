package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateService(saf *models.Saf, model *models.Model) {
	templateName := "service.tmpl"
	pathToDir, err := g.CreateDir(fmt.Sprintf("%s/%s", g.projectRoot, model.Name))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	outputFile, err := g.CreateFile(fmt.Sprintf("%s/%s.service.go", pathToDir, model.Name))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer outputFile.Close()

	rawData := g.fetchServiceRawData(saf, model)
	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func (g *Generator) fetchServiceRawData(saf *models.Saf, model *models.Model) *raw.ServiceRawData {
	var methods []raw.ServiceMethod
	for _, endpoint := range model.Endpoints {
		methodName := "Handle" + endpoint.Type + g.ToUC(model.Name)
		dtoName := fmt.Sprintf("%s%sDTO", g.ToUC(endpoint.Type), g.ToUC(model.Name))

		method := raw.ServiceMethod{
			Type:        endpoint.Type,
			Name:        methodName,
			ResponseDTO: dtoName,
			DTOFields:   endpoint.ResponseDTO,
			Endpoint:    &endpoint,
		}
		methods = append(methods, method)
	}

	rawData := raw.ServiceRawData{
		ModelName:   model.Name,
		ModelNameUC: g.ToUC(model.Name),
		Methods:     methods,
		General:     saf.General,
		Fields:      model.Fields,
	}

	return &rawData
}
