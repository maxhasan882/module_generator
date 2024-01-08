package generator

import (
	"bytes"
	"text/template"
)

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
