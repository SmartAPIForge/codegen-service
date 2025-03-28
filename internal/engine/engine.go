package engine

import (
	"codegen-service/internal/engine/generator"
	"codegen-service/internal/engine/models"
	"encoding/json"
	"fmt"
	"log"
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
		log.Println("Error parsing to saf: ", err)
		return nil, err
	}

	return &saf, nil
}

func (e *Engine) SetupGenerator(saf *models.Saf) {
	log.Println("1.1")
	projectRoot := fmt.Sprintf("output/%s", saf.General.Id)
	log.Println("1.2")
	e.Generator = generator.NewGenerator(projectRoot)
	log.Println("1.3")
	_, err := e.Generator.CreateDir(projectRoot)
	log.Println("1.4")
	if err != nil {
		fmt.Println(err.Error())
		panic("project root setup error")
	}
}

func (e *Engine) Proceed(saf *models.Saf) {
	log.Println("1")
	e.SetupGenerator(saf)

	log.Println("2")
	e.Generator.CopyDockerfile()
	e.Generator.GenerateCompose(saf)
	e.Generator.CopyTaskFile()

	e.Generator.GenerateMigrations(saf)
	e.Generator.GenerateDB()

	e.Generator.GenerateMod(saf)
	e.Generator.GenerateMain(saf)
	log.Println("3")
	// todo generate auth
	for _, model := range saf.Models {
		e.Generator.GenerateModel(&model)
		e.Generator.GenerateService(saf, &model)
		log.Println("3.i")
		e.Generator.GenerateDTOs(&model)
		e.Generator.GenerateController(&model)
	}
}
