package engine

import (
	"codegen-service/internal/engine/generator"
	"codegen-service/internal/engine/models"
	"encoding/json"
	"fmt"
)

type Engine struct {
	source    []byte
	Generator *generator.Generator
}

func NewEngine(contract string) *Engine {
	return &Engine{
		source: []byte(contract),
	}
}

func (e *Engine) ParseSourceToSAF() (*models.Saf, error) {
	var saf models.Saf
	err := json.Unmarshal(e.source, &saf)
	if err != nil {
		return nil, err
	}

	return &saf, nil
}

func (e *Engine) SetupGenerator(saf *models.Saf) {
	projectRoot := fmt.Sprintf("output/%s", saf.General.Id)
	e.Generator = generator.NewGenerator(projectRoot)
	_, err := e.Generator.CreateDir(projectRoot)
	if err != nil {
		panic("project root setup error")
	}
}

func (e *Engine) Proceed(saf *models.Saf) {
	e.SetupGenerator(saf)

	e.Generator.CopyDockerfile()
	e.Generator.GenerateCompose(saf)
	e.Generator.CopyTaskFile()

	e.Generator.GenerateMigrations(saf)
	e.Generator.GenerateDB()

	e.Generator.GenerateMod(saf)
	e.Generator.GenerateMain(saf)
	// todo generate auth
	for _, model := range saf.Models {
		e.Generator.GenerateModel(&model)
		// todo generate service
		e.Generator.GenerateDTOs(&model)
		e.Generator.GenerateController(&model)
	}
}
