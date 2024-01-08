package generator

import (
	"bytes"
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
	var buf bytes.Buffer
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal}).Parse(stubs.GenerateFunc)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, model)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
