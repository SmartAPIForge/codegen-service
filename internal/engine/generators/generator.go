package generators

import "text/template"

type Generator struct {
	templates *template.Template
	outputDir string
}

func NewGenerator() *Generator {
	templates, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	return &Generator{
		templates,
		"output",
	}
}
