package generator

import (
	"codegen-service/internal/engine/models"
	"fmt"
)

func (g *Generator) GenerateMod(saf *models.Saf) {
	templateName := "mod.tmpl"
	pathToDir, err := g.CreateDir(g.projectRoot)
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/go.mod", pathToDir))
	defer outputFile.Close()

	err = g.templates.ExecuteTemplate(outputFile, templateName, saf)
	if err != nil {
		panic(err)
	}
}
