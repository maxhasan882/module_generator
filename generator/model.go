package generator

import (
	"bytes"
	"path/filepath"
	"template/file"
	"text/template"
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

func generateDomainModels(data *Data) {
	err := file.CreateOrUpdateModule(filepath.Join([]string{"domain"}...), data.FileName, "package domain\n")
	if err != nil {
		panic(err)
	}
	for _, model := range data.Models {
		functionData, err := generateFunction(model)
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

func generateSchemaModels(data *Data) {
	err := file.CreateOrUpdateModule(filepath.Join([]string{"infrastructure", "repository", "schema"}...), data.FileName, "package schema")
	if err != nil {
		panic(err)
	}
	for _, model := range data.Models {
		functionData, err := generateGetDataFunction(model)
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
		//err = file.CreateOrUpdateSchemaModule(data.FileName, formattedCode, "schema")
		err = file.CreateOrUpdateModule(filepath.Join([]string{"infrastructure", "repository", "schema"}...), data.FileName, formattedCode)
		if err != nil {
			panic(err)
		}
	}
}

func generateDomainStruct(structName string, fieldsInfo []*FieldInfo) (string, error) {
	data := struct {
		StructName string
		Fields     []*FieldInfo
	}{
		StructName: structName,
		Fields:     fieldsInfo,
	}

	var buf bytes.Buffer
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal, "toLower": toLower}).Parse(domainStructTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generateSchemaStruct(moduleName, structName string, fieldsInfo []*FieldInfo) (string, error) {
	type FieldInfoTemplate struct {
		Name     string
		Type     string
		Optional bool
		JSONTag  string
		IsID     bool
	}

	data := struct {
		ModuleName string
		StructName string
		Fields     []FieldInfoTemplate
		GotID      bool
	}{
		ModuleName: moduleName,
		StructName: structName,
	}

	for _, fieldInfo := range fieldsInfo {
		fieldInfoTemplate := FieldInfoTemplate{
			Name:     fieldInfo.Name,
			Type:     fieldInfo.Type,
			Optional: fieldInfo.Optional,
			IsID:     fieldInfo.Name == "ID" || fieldInfo.Name == "Id",
		}
		if fieldInfoTemplate.IsID {
			data.GotID = true
		}
		data.Fields = append(data.Fields, fieldInfoTemplate)
	}

	var buf bytes.Buffer
	tmpl, err := template.New("schemaStruct").Funcs(template.FuncMap{"snakeToPascal": snakeToPascal, "toLower": toLower}).Parse(schemaStructTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
