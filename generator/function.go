package generator

import (
	"bytes"
	"text/template"
)

func generateSchemasFunction(model *Model) (string, error) {
	var buf bytes.Buffer
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal}).Parse(getDataFunctionTemplate)
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
	tmpl, err := template.New("generateFunc").Funcs(template.FuncMap{"toCamelCase": toCamelCase, "snakeToPascal": snakeToPascal}).Parse(templateText)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&buf, model)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
