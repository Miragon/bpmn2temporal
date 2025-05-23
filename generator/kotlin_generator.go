package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func mapGoToKotlin(goType string) string {
	switch goType {
	case "string":
		return "String"
	case "int":
		return "Int"
	case "bool":
		return "Boolean"
	default:
		return "Any"
	}
}

func GenerateKotlinDataClass(outputPath string, model StructModel) error {
	for i, f := range model.Fields {
		model.Fields[i].Type = mapGoToKotlin(f.Type)
	}
	model.LastIndex = len(model.Fields) - 1

	tmplPath := filepath.Join("templates", "kotlin", "data_class.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, model)
}
