package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
	"io"
	"os"
)

func (g *Generator) GenerateCompose(saf *models.Saf) {
	templateName := "compose.tmpl"
	pathToDir, err := g.CreateDir(g.projectRoot)
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/docker-compose.yml", pathToDir))
	defer outputFile.Close()

	rawData := g.fetchComposeRawData(saf)
	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) fetchComposeRawData(saf *models.Saf) *raw.ComposeRawData {
	rawData := raw.ComposeRawData{
		Port: saf.General.Port,
	}
	return &rawData
}

func (g *Generator) CopyDockerfile() {
	srcPath := "./assets/Dockerfile"
	destPath := fmt.Sprintf("%s/Dockerfile", g.projectRoot)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open source Dockerfile: %v", err))
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create destination Dockerfile: %v", err))
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		panic(fmt.Sprintf("Failed to copy Dockerfile: %v", err))
	}
}

func (g *Generator) CopyTaskFile() {

}
