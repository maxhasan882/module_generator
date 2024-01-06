package generator

import (
	"bytes"
	"strings"
	"text/template"
)

const templateText = `
{{- define "toCamelCase" -}}
	{{ toCamelCase . }}
	{{- end -}}

{{- define "snakeToPascal" -}}
	{{ snakeToPascal . }}
	{{- end -}}

func (m *{{ .ModelName }}) Generate({{ range $index, $field := .Fields }}{{ if $index }}, {{ end }}{{ $field.Name | toCamelCase }} {{ if eq $field.Optional true }}*{{ end }} {{ $field.Type }}{{ end }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ .Name | snakeToPascal  }}: {{ .Name | toCamelCase }},
		{{ end }}
	}
}`

const getDataFunctionTemplate = `
func (m *{{ .ModelName }}) GetData(d *domain.{{ .ModelName }}) *{{ .ModelName }} {
	return &{{ .ModelName }}{
		{{ range .Fields }}{{ if not (or (eq .Name "ID") (eq .Name "Id")) }}{{ .Name | snakeToPascal }}: d.{{ .Name | snakeToPascal }},
		{{ end }}{{ end }}
	}
}
`

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	result := parts[0]
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
