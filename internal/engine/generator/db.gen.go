package generator

import (
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateDB() {
	templateName := "db.tmpl"
	pathToDir, err := g.CreateDir(fmt.Sprintf("%s/db", g.projectRoot))
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/db.go", pathToDir))
	defer outputFile.Close()

	rawData := raw.DBRawData{
		DBPath: fmt.Sprintf("%s/main.db", pathToDir),
	}

	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}
