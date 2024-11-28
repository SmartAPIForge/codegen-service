package generators

import (
	"codegen-service/internal/engine/tree"
	"fmt"
	"os"
)

func (g *Generator) GenerateMain(general *tree.General) {
	templateName := "main.tmpl"

	outputFilePath := fmt.Sprintf("%s/main.go", g.outputDir)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = g.templates.ExecuteTemplate(outputFile, templateName, general)
	if err != nil {
		panic(err)
	}
}
