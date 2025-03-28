package generator

import (
	"codegen-service/internal/engine/models"
	"codegen-service/internal/engine/raw"
	"fmt"
)

func (g *Generator) GenerateMigrations(saf *models.Saf) {
	templateName := "migrations.tmpl"
	pathToDir, err := g.CreateDir(fmt.Sprintf("%s/migrations", g.projectRoot))
	outputFile, err := g.CreateFile(fmt.Sprintf("%s/migration.sql", pathToDir))
	defer outputFile.Close()

	rawData := g.fetchMigrationRawData(saf)
	err = g.templates.ExecuteTemplate(outputFile, templateName, rawData)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) fetchMigrationRawData(saf *models.Saf) *raw.MigrationRawData {
	var extendedModels []raw.ExtendedMigrationModel
	for _, model := range saf.Models {
		var fields []raw.ExtendedMigrationField
		for _, field := range model.Fields {
			extendedField := raw.ExtendedMigrationField{
				Name:     field.Name,
				SQLType:  g.mapTypeToSQL(field.Type),
				IsUnique: field.IsUnique,
			}
			fields = append(fields, extendedField)
		}
		extendedModel := raw.ExtendedMigrationModel{
			NameUC:          g.ToUC(model.Name),
			NameLC:          g.ToLower(model.Name),
			Fields:          fields,
		}
		extendedModels = append(extendedModels, extendedModel)
	}
	rawData := raw.MigrationRawData{
		Models: extendedModels,
	}

	return &rawData
}

func (g *Generator) mapTypeToSQL(goType string) string {
	switch goType {
	case "int":
		return "INTEGER"
	case "string":
		return "TEXT"
	case "bool":
		return "BOOLEAN"
	default:
		return "TEXT"
	}
}
