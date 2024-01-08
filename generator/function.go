package generator

import (
	"bytes"
	"os"
	"template/generator/stubs"
	"text/template"
)

func generateSchemasFunction(model *Model) (string, error) {
	var buf bytes.Buffer
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal}).Parse(stubs.GetDataFunctionTemplate)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, model)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generateDomainsFunction(model *Model) (string, error) {
	fileData, err := os.ReadFile("generator/stubs/generate_func.stub")
	var buf bytes.Buffer
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal}).Parse(string(fileData))
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, model)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
