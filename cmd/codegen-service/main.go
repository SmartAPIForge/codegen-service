package main

import (
	"codegen-service/internal/engine"
	"fmt"
	"io/ioutil"
)

func main() {
	contract, _ := readFile("example.saf.json")

	// setup generation engine
	eng := engine.NewEngine(contract)
	saf := eng.ParseSourceToSAF()
	eng.SetupGenerator(saf)

	// setup project launch environment
	eng.Generator.CopyDockerfile()
	eng.Generator.CopyDockerCompose()
	eng.Generator.CopyTaskFile()

	// setup & generate migrations
	eng.Generator.GenerateMigrations(saf)

	// create db & db client
	eng.Generator.GenerateDB()

	// generate golang deps
	eng.Generator.GenerateMod(saf)

	// generate api
	eng.Generator.GenerateMain(saf)
	// todo generate auth
	for _, model := range saf.Models {
		eng.Generator.GenerateModel(&model)
		// todo generate service
		eng.Generator.GenerateDTOs(&model)
		eng.Generator.GenerateController(&model)
	}
}

func readFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}
