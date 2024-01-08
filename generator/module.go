package generator

import (
	"path/filepath"
	"template/file"
)

// Model represents a data model in the YAML configuration.
type Model struct {
	ModelName string       `yaml:"name"`
	Fields    []*FieldInfo `yaml:"fields"`
}

// FieldInfo represents information about a field in a data model.
type FieldInfo struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Optional bool   `yaml:"optional"`
}

func generateDomainModule(data *Data) {
	err := file.CreateOrUpdateModule(filepath.Join([]string{"domain"}...), data.FileName, "package domain\n")
	if err != nil {
		panic(err)
	}
	for _, model := range data.Models {
		functionData, err := generateDomainsFunction(model)
		if err != nil {
			panic(err)
		}
		generatedStruct, err := generateDomainStruct(model.ModelName, model.Fields)
		if err != nil {
			panic(err)
		}
		formattedCode, err := formatGoCode(generatedStruct + "\n" + functionData)
		if err != nil {
			panic(err)
		}
		err = file.CreateOrUpdateModule(filepath.Join([]string{"domain"}...), data.FileName, formattedCode)
		if err != nil {
			panic(err)
		}
	}
}

func generateInfraModule(data *Data) {
	err := file.CreateOrUpdateModule(filepath.Join([]string{"infrastructure", "repository", "schema"}...), data.FileName, "package schema")
	if err != nil {
		panic(err)
	}
	for _, model := range data.Models {
		functionData, err := generateSchemasFunction(model)
		if err != nil {
			panic(err)
		}
		generatedStruct, err := generateSchemaStruct(data.Module, model.ModelName, model.Fields)
		if err != nil {
			panic(err)
		}
		formattedCode, err := formatGoCode(generatedStruct + functionData)
		if err != nil {
			panic(err)
		}
		err = file.CreateOrUpdateModule(filepath.Join([]string{"infrastructure", "repository", "schema"}...), data.FileName, formattedCode)
		if err != nil {
			panic(err)
		}
	}
}
