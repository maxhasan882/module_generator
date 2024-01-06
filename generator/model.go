package generator

import (
	"fmt"
	"path/filepath"
	"strings"
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

func generateModels(data *Data) {
	err := file.CreateOrUpdateModule(filepath.Join([]string{"domain"}...), data.FileName, "package domain")
	if err != nil {
		panic(err)
	}
	for _, model := range data.Models {
		functionData, err := generateFunction(model)
		if err != nil {
			panic(err)
		}
		generatedStruct := generateDomainStruct(model.ModelName, model.Fields)
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

func generateSchemaModels(data *Data) {
	for _, model := range data.Models {
		functionData, err := generateGetDataFunction(model)
		if err != nil {
			panic(err)
		}
		generatedStruct := generateSchemaStruct(data.Module, model.ModelName, model.Fields)
		formattedCode, err := formatGoCode(generatedStruct + functionData)
		if err != nil {
			panic(err)
		}
		err = file.CreateOrUpdateSchemaModule(data.FileName, formattedCode, "schema")
		if err != nil {
			panic(err)
		}
	}
}

func generateDomainStruct(structName string, fieldsInfo []*FieldInfo) string {
	var structDefinition strings.Builder

	structDefinition.WriteString(fmt.Sprintf("\n\ntype %s struct {\n", structName))

	for _, fieldInfo := range fieldsInfo {
		if fieldInfo.Optional {
			structDefinition.WriteString(fmt.Sprintf("    %s *%s `json:\"%s,omitempty\"`\n", snakeToPascal(fieldInfo.Name), fieldInfo.Type, strings.ToLower(fieldInfo.Name)))
		} else {
			structDefinition.WriteString(fmt.Sprintf("    %s %s `json:\"%s,omitempty\"`\n", snakeToPascal(fieldInfo.Name), fieldInfo.Type, strings.ToLower(fieldInfo.Name)))
		}
	}

	structDefinition.WriteString("}")

	return structDefinition.String()
}

func generateSchemaStruct(moduleName, structName string, fieldsInfo []*FieldInfo) string {
	var structDefinition strings.Builder
	var gotId bool

	structDefinition.WriteString(fmt.Sprintf("\n\ntype %s struct {\n", structName))

	for _, fieldInfo := range fieldsInfo {
		if fieldInfo.Name == "Id" || fieldInfo.Name == "ID" {
			structDefinition.WriteString("Id                primitive.ObjectID `json:\"id,omitempty\" bson:\"_id,omitempty\"`\n")
			gotId = true
			continue
		}
		if fieldInfo.Optional {
			structDefinition.WriteString(fmt.Sprintf("    %s *%s `json:\"%s,omitempty\" bson:\"%s,omitempty\"`\n", snakeToPascal(fieldInfo.Name), fieldInfo.Type, strings.ToLower(fieldInfo.Name), strings.ToLower(fieldInfo.Name)))
		} else {
			structDefinition.WriteString(fmt.Sprintf("    %s %s `json:\"%s,omitempty\" bson:\"%s,omitempty\"`\n", snakeToPascal(fieldInfo.Name), fieldInfo.Type, strings.ToLower(fieldInfo.Name), strings.ToLower(fieldInfo.Name)))
		}
	}

	structDefinition.WriteString("}")
	result := fmt.Sprintf("\n\n"+`import (`+"\n"+`"%s/domain"`+"\n", moduleName)
	if gotId {
		result += `"go.mongodb.org/mongo-driver/bson/primitive"` + "\n"
	}
	result += ")\n"
	return result + structDefinition.String()
}
