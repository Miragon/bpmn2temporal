# BPMN 2 Temporal

## About the project

The CLI tool `b2t` transforms a given BPMN process into boilerplate code of your selected
programming language.

## Development

We use [cobra](https://github.com/spf13/cobra) for creating our CLI tool.  
Additionally, you can install [cobra-cli](https://github.com/spf13/cobra-cli) as a helper to create boilerplate code.

Install `cobra-cli`
```bash
go install github.com/spf13/cobra-cli@latest
```

Add a new command
```bash
cobra-cli add <command>
```

Usage of `b2t`
```bash
go run main.go generate --lang go --output person.go
go run main.go generate --lang kotlin --output Person.kt
```