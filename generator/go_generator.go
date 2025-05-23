package generator

import (
	"os"
	"path/filepath"
	"text/template"
)

func GenerateGoStruct(outputPath string, model StructModel) error {
	model.LastIndex = len(model.Fields) - 1
	tmplPath := filepath.Join("templates", "go", "struct.tmpl")
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
