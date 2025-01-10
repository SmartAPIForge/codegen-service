package engine

import (
	"codegen-service/internal/engine/generator"
	"codegen-service/internal/engine/models"
	"encoding/json"
	"fmt"
	"strconv"
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

func (e *Engine) ParseSourceToSAF() *models.Saf {
	var saf models.Saf
	err := json.Unmarshal(e.source, &saf)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling: %s", err))
	}

	return &saf
}

func (e *Engine) SetupGenerator(saf *models.Saf) {
	projectRoot := strconv.Itoa(saf.General.Id)
	e.Generator = generator.NewGenerator(projectRoot)
	_, err := e.Generator.CreateDir(projectRoot)
	if err != nil {
		panic("project root setup error")
	}

}
