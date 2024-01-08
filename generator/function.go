package generator

import (
	"bytes"
	"strings"
	"text/template"
)

func toLower(s string) string {
	return strings.ToLower(s)
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	result := toLower(parts[0])
	for _, part := range parts[1:] {
		result += strings.Title(part)
	}
	return result
}

func snakeToPascal(snakeCase string) string {
	// Split the string into words based on underscores
	words := strings.Split(snakeCase, "_")

	// Capitalize the first letter of each word
	for i := range words {
		words[i] = strings.Title(words[i])
	}

	// Join the words to form the PascalCase string
	pascalCase := strings.Join(words, "")

	return pascalCase
}

func generateGetDataFunction(model *Model) (string, error) {
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

func generateFunction(model *Model) (string, error) {
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
