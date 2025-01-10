package generator

import (
	"fmt"
	"io"
	"os"
)

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

func (g *Generator) CopyDockerCompose() {
	srcPath := "./assets/docker-compose.yml"
	destPath := fmt.Sprintf("%s/docker-compose.yml", g.projectRoot)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open source docker-compose.yml: %v", err))
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create destination docker-compose.yml: %v", err))
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		panic(fmt.Sprintf("Failed to copy docker-compose.yml: %v", err))
	}
}

func (g *Generator) CopyTaskFile() {
	srcPath := "./assets/Taskfile.yml"
	destPath := fmt.Sprintf("%s/Taskfile.yml", g.projectRoot)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open source Taskfile.yml: %v", err))
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to create destination Taskfile.yml: %v", err))
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		panic(fmt.Sprintf("Failed to copy Taskfile.yml: %v", err))
	}
}
