package cmd

import (
	"fmt"
	"github.com/miragon/b2t/generator"
	"github.com/spf13/cobra"
)

var (
	lang   string
	output string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate boilerplate code",
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	model := generator.StructModel{
		Name: "Person",
		Fields: []generator.Field{
			{"Name", "string"},
			{"Age", "int"},
		},
	}

	var err error
	switch lang {
	case "go":
		err = generator.GenerateGoStruct(output, model)
	case "kotlin":
		err = generator.GenerateKotlinDataClass(output, model)
	default:
		err = fmt.Errorf("given programming language %s is not supported", lang)
	}

	if err != nil {
		fmt.Println("Error while generating code:", err)
	} else {
		fmt.Println("Code was generated under:", output)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&lang, "lang", "l", "go", "Programming Language (go|kotlin)")
	generateCmd.Flags().StringVarP(&output, "output", "o", "output.txt", "Output path")

	rootCmd.AddCommand(generateCmd)
}
