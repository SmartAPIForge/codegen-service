package engine

import (
	"codegen-service/internal/engine/generators"
	"codegen-service/internal/engine/tree"
	"encoding/json"
	"fmt"
)

type Engine struct {
	source    []byte
	Generator *generators.Generator
}

func NewEngine(contract string) *Engine {
	source := []byte(contract)
	generator := generators.NewGenerator()

	return &Engine{
		source,
		generator,
	}
}

func (e *Engine) ParseSourceToSAF() *tree.Saf {
	var saf tree.Saf
	err := json.Unmarshal(e.source, &saf)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling: %s", err))
	}

	return &saf
}
