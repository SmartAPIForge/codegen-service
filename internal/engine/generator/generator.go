package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Generator struct {
	templates   *template.Template
	projectRoot string
}

func NewGenerator(projectRoot string) *Generator {
	// Создаем функции для шаблонизатора
	funcMap := template.FuncMap{
		"ToUC": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToUpper(string(s[0])) + s[1:]
		},
	}
	
	// Создаем шаблонизатор с функциями
	templates := template.New("").Funcs(funcMap)
	
	// Парсим все шаблоны
	templates, err := templates.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	return &Generator{
		templates,
		projectRoot,
	}
}

func (*Generator) CreateDir(path string) (string, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	return path, nil
}

func (*Generator) CreateFile(outputFilePath string) (*os.File, error) {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	return outputFile, nil
}

func (*Generator) ToUC(text string) string {
	if len(text) == 0 {
		return text
	}
	text = strings.ToLower(text)
	return strings.ToUpper(string(text[0])) + text[1:]
}

func (*Generator) ToLower(text string) string {
	if len(text) == 0 {
		return text
	}
	return strings.ToLower(text)
}
